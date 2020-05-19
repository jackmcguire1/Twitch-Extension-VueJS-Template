package twitch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func intFromHeader(name string, header *http.Header) (i int) {
	v := header.Get(name)
	i, _ = strconv.Atoi(v)
	return
}

func (t *Twitch) do(
	method string,
	url string,
	b []byte,
	q url.Values,
	auth string,
	authType string,
) (
	data []byte,
	headers http.Header,
	err error,
) {
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		err = fmt.Errorf("failed to construct request err:%s", err)
		return
	}
	err = t.setExtensionRequestHeaders(req, auth, authType)
	if err != nil {
		return
	}
	req.URL.RawQuery = q.Encode()

	resp, err := t.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	headers = resp.Header

	switch resp.StatusCode {
	case http.StatusOK:
		data, err = ioutil.ReadAll(resp.Body)
	case http.StatusNoContent:
	default:
		err = fmt.Errorf(
			"unsupported response httpCode:%d status:%s ",
			resp.StatusCode,
			resp.Status,
		)
		return
	}

	return
}

func (t *Twitch) setExtensionRequestHeaders(
	req *http.Request,
	token string,
	authType string,
) (
	err error,
) {
	if authType == "" {
		authType = "Bearer"
	}

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", authType, token))
	}
	req.Header.Set("Client-ID", t.ClientId)
	req.Header.Set("Content-Type", "application/json")

	return
}
