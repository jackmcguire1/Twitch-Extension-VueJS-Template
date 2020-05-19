package twitch

import (
	"fmt"
	"log"
	"time"

	"github.com/nicklaw5/helix"
)

func (t *Twitch) GetUsers(
	ids []string,
	logins []string,
) (
	users []helix.User,
	err error,
) {
	params := &helix.UsersParams{
		IDs:    ids,
		Logins: logins,
	}

	users, err = t.getUsers(params)
	return
}

func (t *Twitch) GetUser(
	userID string,
) (
	twitchUser helix.User,
	err error,
) {
	params := &helix.UsersParams{
		IDs:    []string{userID},
		Logins: []string{},
	}

	users, err := t.getUsers(params)
	if err != nil {
		return
	}

	if len(users) == 0 {
		err = fmt.Errorf("failed to find user %q", userID)
		return
	}

	twitchUser = users[0]

	return
}

func (t *Twitch) getUsers(params *helix.UsersParams) (users []helix.User, err error) {

	if len(params.Logins)+len(params.IDs) > 100 {
		err = fmt.Errorf("to many get user ids or logins params")
		return
	}

	rsp, err := t.helix.GetUsers(params)
	if err != nil {
		return
	}

	reset := time.Unix(int64(rsp.GetRateLimitReset()), 0).Format(time.RFC3339)
	log.Printf(
		"GetUser rate limiting info rateLimit:%d remaining:%d reset:%q epoch:%d",
		rsp.GetRateLimit(),
		rsp.GetRateLimitRemaining(),
		reset,
		rsp.GetRateLimitReset(),
	)

	if rsp.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			rsp.Error,
			rsp.ErrorMessage,
			rsp.StatusCode,
		)
		return
	}

	users = rsp.Data.Users

	return
}
