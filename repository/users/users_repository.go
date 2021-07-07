package repository

import (
	"bookstore_oauth-api/domain/users"
	resp "bookstore_oauth-api/utils/response"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8000",
		Timeout: 100 * time.Millisecond,
	}
)

type IUserRepository interface {
	Login(*users.UserLogin) (*users.User, *resp.RestErr)
	New() IUserRepository
}

type userRepository struct{}

var UserRepository IUserRepository = &userRepository{}

func (repo *userRepository) Login(ul *users.UserLogin) (*users.User, *resp.RestErr) {
	res := restClient.Post("/users/login", ul)
	if res == nil || res.Response == nil {
		return nil, resp.InternalServerError("failed to login")
	}

	if res.StatusCode > 299 {
		var restErr *resp.RestErr
		err := json.Unmarshal(res.Bytes(), &restErr)
		if err != nil {
			return nil, resp.InternalServerError("failed to parse data response from users service")
		}
		return nil, restErr
	}

	var restResp struct {
		Data *users.User
	}
	if err := json.Unmarshal(res.Bytes(), &restResp); err != nil {
		fmt.Println(err.Error())
		return nil, resp.InternalServerError("failed to parse users data")
	}

	return restResp.Data, nil
}

func (repo *userRepository) New() IUserRepository {
	return &userRepository{}
}
