package request

import (
	"errors"
	"net/url"
	"strconv"
)

type APIRequest struct {
	path  string
	appID string
	token string
	ts    int64
}

func NewAPIRequestFromRawUrl(rawURL string) (*APIRequest, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.New("parse rawUrl error: " + err.Error())
	}
	qs := u.Query()
	timestamp, err := strconv.ParseInt(qs.Get("ts"), 10, 64)
	if err != nil {
		return nil, err
	}
	ar := &APIRequest{
		path:  u.Path,
		appID: qs.Get("appid"),
		token: qs.Get("token"),
		ts:    timestamp,
	}
	return ar, nil
}

func (ar *APIRequest) GetBaseUrl() string {
	return ar.path
}

func (ar *APIRequest) GetAppID() string {
	return ar.appID
}

func (ar *APIRequest) GetToken() string {
	return ar.token
}

func (ar *APIRequest) GetTS() int64 {
	return ar.ts
}
