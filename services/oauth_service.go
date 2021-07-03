package services

import (
	oauth "bookstore_oauth-api/domain/oauth"
	repo "bookstore_oauth-api/repository"
	resp "bookstore_oauth-api/utils/response"
)

type Service interface {
	GetById(string) (*oauth.AccessToken, *resp.RestErr)
}

type service struct {
	repository repo.Repository
}

func (s *service) GetById(tokenId string) (*oauth.AccessToken, *resp.RestErr) {
	at, err := s.repository.GetById(tokenId)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func New(repo repo.Repository) Service {
	return &service{
		repository: repo,
	}
}
