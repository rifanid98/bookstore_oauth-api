package handler

import (
	"bookstore_oauth-api/domain/oauth"
	"bookstore_oauth-api/services"
	resp "bookstore_oauth-api/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
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

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at oauth.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, resp.BadRequest("invalid json body"))
		return
	}
	if err := h.service.Create(&at); err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(at))
	return
}

func (h *accessTokenHandler) Update(c *gin.Context) {
	var at oauth.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, resp.BadRequest("invalid json body"))
		return
	}
	if err := h.service.Update(&at); err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}
	c.JSON(http.StatusOK, resp.Success(at))
	return
}

func New(service services.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
