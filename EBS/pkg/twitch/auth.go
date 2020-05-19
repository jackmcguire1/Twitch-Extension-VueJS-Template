package twitch

import (
	"EBS/m/v2/env"
	"EBS/m/v2/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/jackmcguire1/go-twitch-ext"
	"github.com/nicklaw5/helix"
	"golang.org/x/oauth2"
)

const (
	AUTH_URL = "https://id.twitch.tv/oauth2"

	ExtAnalyticsScope string = "analytics:read:extensions"
	BitsRead          string = "bits:read"
	UserSubscriptions string = "channel:read:subscriptions"
)

type OIDCAuth struct {
	oauth2.Token
	RawIDToken string        `json:"raw_id_token"`
	IDToken    *oidc.IDToken `json:"id_token"`
	Claims     *OIDCClaims   `json:"claims"`
}

type OIDCClaims struct {
	IDToken struct {
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	} `json:"id_token"`
	UserInfo struct {
		Username string `json:"preferred_username"`
		Picture  string `json:"picture"`
		Updated  string `json:"updated_at"`
	} `json:"userinfo"`
}

type TokenValidation struct {
	ClientID string   `json:"client_id"`
	Username string   `json:"login"`
	Scopes   []string `json:"scopes"`
	UserID   string   `json:"user_id"`
}

func (t *Twitch) SetUserAccessToken(code string) {
	t.helix.SetUserAccessToken(code)
}

func (t *Twitch) SetAppAccessToken(code string) {
	t.helix.SetAppAccessToken(code)
}

func (t *Twitch) GetAppAccessToken() (data helix.AppAccessCredentials, err error) {
	resp, err := t.helix.GetAppAccessToken()
	if err != nil {
		return
	}

	reset := time.Unix(int64(resp.GetRateLimitReset()), 0).Format(time.RFC3339)
	log.Printf(
		"GetAppAccessToken rate Limiting info rateLimit:%d remaining:%d reset:%q epoch:%d",
		resp.GetRateLimit(),
		resp.GetRateLimitRemaining(),
		reset,
		resp.GetRateLimitReset(),
	)

	if resp.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			resp.Error,
			resp.ErrorMessage,
			resp.StatusCode,
		)
		return
	}
	data = resp.Data

	return
}

func (t *Twitch) GetAuthURL(scopes []string) (url string) {
	t.helix.SetScopes(scopes)
	url = t.helix.GetAuthorizationURL("", false)

	return
}

func (t *Twitch) ValidateToken(
	ctx context.Context,
	accessToken string,
) (
	validation *TokenValidation,
	err error,
) {
	uri := "https://id.twitch.tv/oauth2/validate"

	resp, _, err := t.do(http.MethodGet, uri, nil, nil, accessToken, "OAuth")
	if err != nil {
		return
	}

	validation = &TokenValidation{}
	err = json.Unmarshal(resp, &validation)

	return
}

func (t *Twitch) UserInfo(
	ctx context.Context,
	token *oauth2.Token,
	v interface{},
) (
	err error,
) {
	provider, err := oidc.NewProvider(ctx, AUTH_URL)
	if err != nil {
		return
	}

	userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
	if err != nil {
		return
	}

	err = userInfo.Claims(&v)
	return
}

func (t *Twitch) UserOIDCAuth(
	ctx context.Context,
	redirectUrl string,
	code string,
) (
	resp *OIDCAuth,
	err error,
) {
	resp = &OIDCAuth{}
	provider, err := oidc.NewProvider(ctx, AUTH_URL)
	if err != nil {
		return
	}

	oauth2Config := oauth2.Config{
		ClientID:     t.ClientId,
		ClientSecret: t.ClientSecret,
		RedirectURL:  redirectUrl,
		Endpoint:     provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: env.ClientId})

	oauth2Token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		return
	}
	resp.Token = *oauth2Token

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		err = fmt.Errorf("id_token is MISSING validate twitch auth req")
		return
	}
	resp.RawIDToken = rawIDToken

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		err = fmt.Errorf("failed to verify id_token err:%s", err)
		return
	}
	resp.IDToken = idToken

	claims := &OIDCClaims{}
	err = idToken.Claims(&claims.IDToken)
	if err != nil {
		return
	}
	resp.Claims = claims

	return
}

func (t *Twitch) GetUserAccessToken(code string, scopes []string) (data helix.UserAccessCredentials, err error) {
	t.helix.SetScopes(scopes)

	resp, err := t.helix.GetUserAccessToken(code)
	if err != nil {
		return
	}

	reset := time.Unix(int64(resp.GetRateLimitReset()), 0).Format(time.RFC3339)
	log.Printf(
		"GetUserAccessToken rate Limiting info rateLimit:%d remaining:%d reset:%q epoch:%d",
		resp.GetRateLimit(),
		resp.GetRateLimitRemaining(),
		reset,
		resp.GetRateLimitReset(),
	)

	if resp.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			resp.Error,
			resp.ErrorMessage,
			resp.StatusCode,
		)
		return
	}
	data = resp.Data

	return
}

func (t *Twitch) RenewUserAccessToken(refreshToken string) (data helix.UserAccessCredentials, err error) {
	resp, err := t.helix.RefreshUserAccessToken(refreshToken)
	if err != nil {
		return
	}

	reset := time.Unix(int64(resp.GetRateLimitReset()), 0).Format(time.RFC3339)
	log.Printf(
		"RenewUserAccessToken rate Limiting info rateLimit:%d remaining:%d reset:%q epoch:%d",
		resp.GetRateLimit(),
		resp.GetRateLimitRemaining(),
		reset,
		resp.GetRateLimitReset(),
	)

	if resp.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			resp.Error,
			resp.ErrorMessage,
			resp.StatusCode,
		)
		return
	}
	data = resp.Data

	return
}

func ExtractExtClaimsFromHeader(
	context map[string]interface{},
) (
	claims *twitchext.TwitchJWTClaims,
	err error,
) {

	claimsStr, ok := context["claims"]
	if !ok {
		err = fmt.Errorf("failed to get claims from auth context")
		return
	}

	claims = &twitchext.TwitchJWTClaims{}
	err = json.Unmarshal([]byte(claimsStr.(string)), claims)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal claims err:%s", err)
		return
	}
	log.Println(utils.ToJSON(claims))

	return
}

func ValidateExtAuthToken(
	headers map[string]string,
	twitchPkg *twitchext.Twitch,
) (
	claims *twitchext.TwitchJWTClaims,
	err error,
) {
	authHeader, ok := headers["Authorization"]
	if !ok {
		err = fmt.Errorf("Missing Authorization header")
		return
	}

	tokenStr := strings.Split(authHeader, "Bearer ")
	if len(tokenStr) != 2 {
		err = fmt.Errorf("invalid auth header %q", authHeader)
		return
	}

	claims, err = twitchPkg.JWTVerify(tokenStr[1])
	if err != nil {
		return
	}
	log.Println(utils.ToJSON(claims))

	return
}
