package services

import (
	oauth "bookstore_oauth-api/domain/oauth"
	"bookstore_oauth-api/domain/users"
	oauthRepo "bookstore_oauth-api/repository/oauth"
	usersRepo "bookstore_oauth-api/repository/users"
	resp "bookstore_oauth-api/utils/response"
)

type Service interface {
	GetById(string) (*oauth.AccessToken, *resp.RestErr)
	Create(*oauth.AccessTokenRequest) (*oauth.AccessToken, *resp.RestErr)
	Update(*oauth.AccessToken) *resp.RestErr
}

type service struct {
	oauthRepository oauthRepo.IOAuthRepository
	usersRepository usersRepo.IUserRepository
}

func (s *service) GetById(tokenId string) (*oauth.AccessToken, *resp.RestErr) {
	at, err := s.oauthRepository.GetById(tokenId)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *service) Create(atr *oauth.AccessTokenRequest) (*oauth.AccessToken, *resp.RestErr) {
	if err := atr.Validate(); err != nil {
		return nil, err
	}

	user, err := s.usersRepository.Login(&users.UserLogin{
		Username: atr.Username,
		Password: atr.Password,
	})
	if err != nil {
		return nil, err
	}

	if user.Id < 1 {
		return nil, resp.Unauthorized("")
	}

	at := oauth.GetAccessToken(user.Id)
	at.ClientId = atr.ClientId
	at.Generate()

	if err := s.oauthRepository.Create(at); err != nil {
		return nil, err
	}

	return at, nil
}

func (s *service) Update(atr *oauth.AccessToken) *resp.RestErr {
	if err := atr.Validate(); err != nil {
		return err
	}
	return s.oauthRepository.Update(atr)
}

func New(oauthRepo oauthRepo.IOAuthRepository, usersRepo usersRepo.IUserRepository) Service {
	return &service{
		oauthRepository: oauthRepo,
		usersRepository: usersRepo,
	}
}
