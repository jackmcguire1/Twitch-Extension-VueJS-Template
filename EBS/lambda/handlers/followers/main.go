package main

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	log "github.com/apex/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	twitchext "github.com/jackmcguire1/go-twitch-ext"

	"EBS/m/v2/env"
	"EBS/m/v2/pkg/twitch"
	"EBS/m/v2/pkg/utils"
)

type FollowerStore struct {
	Username   string
	Followers  int
	expiration time.Time
}

type FollowersResponse struct {
	Followers int `json:"followers"`
}

var (
	twitchPkg      *twitchext.Twitch
	twitchInternal *twitch.Twitch
	followersCache map[string]*FollowerStore
)

func init() {

	log.SetLevelFromString(env.DebugLevel)

	twitchPkg = twitchext.NewClient(
		env.OwnerID,
		env.ClientId,
		env.ClientSecret,
		env.ExtVersion,
		env.ExtConfigVersion,
	)
	twitchInternal, _ = twitch.New(env.ClientId, env.ClientSecret)

	data, err := twitchInternal.GetAppAccessToken()
	if err != nil {
		panic(err)
	}
	twitchInternal.SetAppAccessToken(data.AccessToken)

	log.
		WithField("access-token-resp", utils.ToJSON(data)).
		Info("got new app access token")

	followersCache = map[string]*FollowerStore{}
}

func handler(
	ctx context.Context,
	event events.APIGatewayProxyRequest,
) (
	resp *events.APIGatewayProxyResponse,
	hardError error,
) {
	log.
		WithField("context", ctx).
		WithField("raw-event", utils.ToJSON(event)).
		Info("got request")

	resp = &events.APIGatewayProxyResponse{}

	resp.Headers = map[string]string{
		"Access-Control-Allow-Origin":      event.Headers["origin"],
		"Access-Control-Allow-Methods":     "OPTIONS,GET",
		"Access-Control-Allow-Headers":     "Content-Type,Authorization,X-Requested-With,Origin,Accept",
		"Access-Control-Allow-Credentials": "true",
		"Content-Type":                     "application/json",
	}
	resp.StatusCode = http.StatusOK

	switch event.HTTPMethod {
	case http.MethodOptions:
		log.
			WithField("context", ctx).
			WithField("raw-event", utils.ToJSON(event)).
			WithField("raw-response", utils.ToJSON(resp)).
			Info("found options event")

		return
	case http.MethodGet:
	default:
		resp.StatusCode = http.StatusBadRequest

		log.
			WithField("context", ctx).
			WithField("raw-event", utils.ToJSON(event)).
			WithField("raw-response", utils.ToJSON(resp)).
			Warn("invalid http request")

		return
	}
	userID := filepath.Base(event.Path)

	log.
		WithField("context", ctx).
		WithField("raw-event", utils.ToJSON(event)).
		WithField("userID", userID).
		Info("got request")

	claims, err := twitch.ValidateExtAuthToken(event.Headers, twitchPkg)
	if err != nil {
		log.
			WithError(err).
			WithField("context", ctx).
			WithField("raw-event", utils.ToJSON(event)).
			WithField("userID", userID).
			Warn("failed to retrieve claims for twitch extension user")
	}

	entryLog := log.
		WithField("context", ctx).
		WithField("raw-event", utils.ToJSON(event)).
		WithField("userID", userID).
		WithField("claims", utils.ToJSON(claims))

	entryLog.Info("got claims from authorization header")

	userFollowers, ok := followersCache[userID]
	if ok {
		if time.Now().Before(userFollowers.expiration) {
			resp.Body = utils.ToJSON(&FollowersResponse{Followers: userFollowers.Followers})
			entryLog.
				WithField("raw-resp", utils.ToJSON(resp)).
				WithField("follower", utils.ToJSON(userFollowers)).
				Info("got followers from in-memory cache")

			return
		}
	}

	followersResp, err := twitchInternal.GetChannelFollowers(userID, "", 200)
	if err != nil {
		resp.Body = utils.ToJSON(struct{ Error string }{Error: "failed to get followers"})
		resp.StatusCode = http.StatusInternalServerError

		entryLog.
			WithField("raw-resp", utils.ToJSON(resp)).
			WithError(err).
			Error("failed to get followers from Twitch API")

		return
	}
	followers := followersResp.Data.Total

	followersCache[userID] = &FollowerStore{
		Followers:  followers,
		expiration: time.Now().Add(time.Minute),
	}
	resp.Body = utils.ToJSON(&FollowersResponse{Followers: followers})

	entryLog.
		WithField("raw-resp", utils.ToJSON(resp)).
		WithField("followers-resp", utils.ToJSON(followersResp)).
		Info("got followers from Twitch API")

	return
}

func main() {
	lambda.Start(handler)
}
