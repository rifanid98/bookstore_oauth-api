package oauth

import (
	resp "bookstore_oauth-api/utils/response"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *resp.RestErr {
	accessToken := strings.TrimSpace(at.AccessToken)
	if len(accessToken) < 1 {
		return resp.BadRequest("Access tokens cannot be empty")
	}
	if at.UserId < 1 {
		return resp.BadRequest("User ID cannot be 0")
	}
	if at.ClientId < 1 {
		return resp.BadRequest("Client ID cannot be 0")
	}
	if at.Expires < 1 {
		return resp.BadRequest("Expires cannot be 0")
	}
	return nil
}
