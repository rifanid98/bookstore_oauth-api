package app

import (
	"bookstore_oauth-api/client/cassandra"
	"bookstore_oauth-api/handler"
	"bookstore_oauth-api/repository"
	"bookstore_oauth-api/services"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()

	handler := handler.New(services.New(repository.New()))

	router.GET("/oauth/token/:token_id", handler.GetById)
	router.POST("/oauth/token", handler.Create)
	router.PATCH("/oauth/token", handler.Update)
	router.Run(":8080")
}
