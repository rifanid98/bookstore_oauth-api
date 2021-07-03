package app

import (
	"bookstore_oauth-api/handler"
	"bookstore_oauth-api/repository"
	"bookstore_oauth-api/services"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	handler := handler.New(services.New(repository.New()))

	router.GET("/oauth/token/:token_id", handler.GetById)
	router.POST("/oauth/token", handler.Create)
	router.PATCH("/oauth/token", handler.Update)
	router.Run(":8080")
}
