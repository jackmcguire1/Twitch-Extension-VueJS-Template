package twitch

import (
	"fmt"
	"log"
	"time"

	"github.com/nicklaw5/helix"
)

func (t *Twitch) GetChannelFollowers(
	userID string,
	bookmark string,
) (
	users *helix.UsersFollowsResponse,
	err error,
) {
	params := &helix.UsersFollowsParams{
		ToID:  userID,
		First: 100,
	}

	if bookmark != "" {
		params.After = bookmark
	}

	users, err = t.helix.GetUsersFollows(params)
	if err != nil {
		return
	}

	reset := time.Unix(int64(users.GetRateLimitReset()), 0).Format(time.RFC3339)
	log.Printf(
		"GetChannelFollowers rate Limiting info rateLimit:%d remaining:%d reset:%q epoch:%d",
		users.GetRateLimit(),
		users.GetRateLimitRemaining(),
		reset,
		users.GetRateLimitReset(),
	)

	if users.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			users.Error,
			users.ErrorMessage,
			users.StatusCode,
		)
		return
	}

	return
}

func (t *Twitch) GetChannelsFollows(
	userID string,
	bookmark string,
	limit int,
) (
	users *helix.UsersFollowsResponse,
	err error,
) {
	params := &helix.UsersFollowsParams{
		FromID: userID,
		First:  limit,
	}

	if bookmark != "" {
		params.After = bookmark
	}
	users, err = t.helix.GetUsersFollows(params)
	if err != nil {
		return
	}

	reset := time.Unix(int64(users.GetRateLimitReset()), 0).Format(time.RFC3339)
	log.Printf(
		"GetChannelsFollows rate Limiting info rateLimit:%d remaining:%d reset:%q epoch:%d",
		users.GetRateLimit(),
		users.GetRateLimitRemaining(),
		reset,
		users.GetRateLimitReset(),
	)

	if users.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			users.Error,
			users.ErrorMessage,
			users.StatusCode,
		)
		return
	}

	return
}

func (t *Twitch) FollowsUser(
	toID string,
	fromID string,
) (
	follower *helix.UserFollow,
	following bool,
	err error,
) {
	params := &helix.UsersFollowsParams{
		ToID:   toID,
		FromID: fromID,
	}

	resp, err := t.helix.GetUsersFollows(params)
	if err != nil {
		return
	}

	reset := time.Unix(int64(resp.GetRateLimitReset()), 0).Format(time.RFC3339)
	log.Printf(
		"FollowsUser rate Limiting info rateLimit:%d remaining:%d reset:%q epoch:%d",
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

	if resp.Data.Total == 0 {
		return
	}

	following = true
	follower = &resp.Data.Follows[0]

	return
}
