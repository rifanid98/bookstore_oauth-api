package oauth

import (
	"fmt"
	"strings"
	"time"

	hash "github.com/rifanid98/bookstore_utils-go/hash"
	resp "github.com/rifanid98/bookstore_utils-go/response"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetAccessToken(userId int64) *AccessToken {
	return &AccessToken{
		UserId:  userId,
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

func (at *AccessToken) Generate() {
	at.AccessToken = hash.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	Username string `json:"username"`
	Password string `json:"password"`

	ClientId     int64  `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *resp.RestErr {
	switch at.GrantType {
	case grantTypePassword:
	case grandTypeClientCredentials:
		break

	default:
		return resp.BadRequest("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}
