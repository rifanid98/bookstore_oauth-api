package repository

import (
	"bookstore_oauth-api/client/cassandra"
	oauth "bookstore_oauth-api/domain/oauth"
	resp "bookstore_oauth-api/utils/response"
	"fmt"

	"github.com/gocql/gocql"
)

type IOAuthRepository interface {
	GetById(string) (*oauth.AccessToken, *resp.RestErr)
	Create(*oauth.AccessToken) *resp.RestErr
	Update(*oauth.AccessToken) *resp.RestErr
	New() IOAuthRepository
}

type oauthRepository struct{}

var OAuthRepository IOAuthRepository = &oauthRepository{}

func (repo *oauthRepository) GetById(tokenId string) (*oauth.AccessToken, *resp.RestErr) {
	session := cassandra.GetSession()

	var accessToken oauth.AccessToken
	err := session.Query(GetAccessTokenQuery, tokenId).Scan(
		&accessToken.AccessToken,
		&accessToken.UserId,
		&accessToken.ClientId,
		&accessToken.Expires,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, resp.NotFound(fmt.Sprintf("access token with token id %s not found", tokenId))
		}
		return nil, resp.InternalServerError(err.Error())
	}

	return &accessToken, nil
}

func (repo *oauthRepository) Create(at *oauth.AccessToken) *resp.RestErr {
	session := cassandra.GetSession()

	err := session.Query(
		CreateAccessTokenQuery,
		&at.AccessToken,
		&at.UserId,
		&at.ClientId,
		&at.Expires,
	).Exec()
	if err != nil {
		return resp.InternalServerError(err.Error())
	}

	return nil
}

func (repo *oauthRepository) Update(at *oauth.AccessToken) *resp.RestErr {
	session := cassandra.GetSession()

	err := session.Query(
		UpdateQuery,
		at.UserId,
		at.ClientId,
		at.Expires,
		at.AccessToken,
	).Exec()
	if err != nil {
		return resp.InternalServerError(err.Error())
	}

	return nil
}

func (repo *oauthRepository) New() IOAuthRepository {
	return &oauthRepository{}
}
