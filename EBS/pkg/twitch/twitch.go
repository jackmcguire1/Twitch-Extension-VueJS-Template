package twitch

import (
	"EBS/m/v2/env"
	"github.com/nicklaw5/helix"
	"net/http"
)

type Twitch struct {
	client          http.Client
	AppAccessToken  string
	UserAccessToken string
	ClientId        string
	ClientSecret    string
	helix           *helix.Client
}

func New(clientID string, clientSecret string) (t *Twitch, err error) {

	client, err := helix.NewClient(&helix.Options{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		AppAccessToken: env.AppAccessToken,
		RedirectURI:    env.TwitchAuthRedirect,
	})
	if err != nil {
		return
	}

	t = &Twitch{
		ClientId:       clientID,
		ClientSecret:   clientSecret,
		AppAccessToken: env.AppAccessToken,
		helix:          client,
	}
	return
}
