package repository

import (
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	email    = os.Getenv("EMAIL_TEST")
	password = os.Getenv("PASSWORD_TEST")
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {

}

func TestLoginUserFailedToLogin(t *testing.T) {

}

func TestLoginUserInvalidCredentials(t *testing.T) {

}

func TestLoginUserInvalidJsonResponse(t *testing.T) {

}
