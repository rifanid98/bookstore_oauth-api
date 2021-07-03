package repository

import (
	"bookstore_oauth-api/client/cassandra"
	oauth "bookstore_oauth-api/domain/oauth"
	resp "bookstore_oauth-api/utils/response"
)

type Repository interface {
	GetById(string) (*oauth.AccessToken, *resp.RestErr)
}

type repository struct{}

func (repo *repository) GetById(tokenId string) (*oauth.AccessToken, *resp.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, resp.InternalServerError("The database has not been implemented")
}

func New() Repository {
	return &repository{}
}
