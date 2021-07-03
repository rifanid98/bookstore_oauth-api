package handler

import (
	"bookstore_oauth-api/services"
	resp "bookstore_oauth-api/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service services.Service
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	token, err := h.service.GetById(strings.TrimSpace(c.Param("token_id")))
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(token))
}

func New(service services.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
