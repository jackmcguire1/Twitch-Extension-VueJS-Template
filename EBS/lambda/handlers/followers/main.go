package main

import (
	"context"
	log "github.com/apex/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	twitchext "github.com/jackmcguire1/go-twitch-ext"
	"net/http"
	"path/filepath"

	"EBS/m/v2/env"
	"EBS/m/v2/pkg/twitch"
	"EBS/m/v2/pkg/utils"
)

type FollowersResponse struct {
	Total     int         `json:"total"`
	Followers []*Follower `json:"followers"`
}

type Follower struct {
	Username string    `json:"username"`
}

var (
	twitchPkg      *twitchext.Twitch
	twitchInternal *twitch.Twitch
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
		"Access-Control-Allow-Origin":      event.Headers["Origin"],
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

	var total int
	var followers []*Follower
	var requestCount int64
	var bookmark string
	for {
		log.
			WithField("request-num", requestCount).
			WithField("userID", userID).
			Info("getting followers from Twitch API")

		followersResp, err := twitchInternal.GetChannelFollowers(userID, bookmark)
		if err != nil {
			resp.Body = utils.ToJSON(struct{ Error string }{Error: "failed to get followers"})
			resp.StatusCode = http.StatusInternalServerError

			entryLog.
				WithField("raw-resp", utils.ToJSON(resp)).
				WithError(err).
				Error("failed to get followers from Twitch API")

			return
		}
		log.
			WithField("request-num", requestCount).
			WithField("userID", userID).
			WithField("twitch-response", utils.ToJSON(followersResp)).
			Debug("got followers from Twitch API")

		total = followersResp.Data.Total

		for _, follower := range followersResp.Data.Follows {
			followers = append(followers, &Follower{
				Username: follower.FromName,
			})
		}

		if followersResp.Data.Pagination.Cursor == "" {
			log.
				WithField("request-num", requestCount).
				WithField("userID", userID).
				WithField("twitch-response", utils.ToJSON(followersResp)).
				Debug("no cursor found, finished polling for followers")

			break
		}
		bookmark = followersResp.Data.Pagination.Cursor

		requestCount++
	}
	followersResp := utils.ToJSON(&FollowersResponse{Followers: followers, Total: total})
	resp.Body = followersResp

	entryLog.
		WithField("followers-resp", followersResp).
		WithField("raw-resp", utils.ToJSON(resp)).
		Info("returning followers from Twitch API")

	return
}

func main() {
	lambda.Start(handler)
}
