package repository

import (
	"bookstore_oauth-api/client/cassandra"
	oauth "bookstore_oauth-api/domain/oauth"
	resp "bookstore_oauth-api/utils/response"
	"fmt"

	"github.com/gocql/gocql"
)

type Repository interface {
	GetById(string) (*oauth.AccessToken, *resp.RestErr)
	Create(*oauth.AccessToken) *resp.RestErr
	Update(*oauth.AccessToken) *resp.RestErr
}

type repository struct{}

func (repo *repository) GetById(tokenId string) (*oauth.AccessToken, *resp.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, resp.InternalServerError(err.Error())
	}
	defer session.Close()

	var accessToken oauth.AccessToken
	err = session.Query(GetAccessTokenQuery, tokenId).Scan(
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

func (repo *repository) Create(at *oauth.AccessToken) *resp.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return resp.InternalServerError(err.Error())
	}
	defer session.Close()

	err = session.Query(
		CreateAccessTokenQuery,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec()
	if err != nil {
		return resp.InternalServerError(err.Error())
	}

	return nil
}

func (repo *repository) Update(at *oauth.AccessToken) *resp.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return resp.InternalServerError(err.Error())
	}
	defer session.Close()

	err = session.Query(
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

func New() Repository {
	return &repository{}
}
