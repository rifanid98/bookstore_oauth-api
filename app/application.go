package app

import (
	"bookstore_oauth-api/handler"
	oauthRepository "bookstore_oauth-api/repository/oauth"
	usersRepository "bookstore_oauth-api/repository/users"
	"bookstore_oauth-api/services"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	oauthRepo := oauthRepository.OAuthRepository.New()
	usersRepo := usersRepository.UserRepository.New()
	handler := handler.New(services.New(oauthRepo, usersRepo))

	router.GET("/oauth/token/:token_id", handler.GetById)
	router.POST("/oauth/token", handler.Create)
	router.PATCH("/oauth/token", handler.Update)
	router.Run(":8001")
}
