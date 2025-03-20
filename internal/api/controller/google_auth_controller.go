package controller

import (
	"net/http"

	"github.com/Matheus-Lara/orare/internal/api"
	"github.com/Matheus-Lara/orare/internal/api/service"
	"github.com/gin-gonic/gin"
)

type GoogleAuthController struct {
	service *service.GoogleAuthService
}

func (controller *GoogleAuthController) StartGoogleAuthorization(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, controller.service.BuildGoogleAuthURL())
}

func (controller *GoogleAuthController) AuthorizationCode(c *gin.Context) {
	err := controller.service.ExchangeCodeForToken(c.Query("code"), c.Query("state"))

	if err != nil {
		api.ResponseBadRequest(c, err)
		return
	}

	api.ResponseSuccess(c, http.StatusOK, gin.H{"message": "success"})
}

func NewGoogleAuthController(service *service.GoogleAuthService) *GoogleAuthController {
	return &GoogleAuthController{
		service: service,
	}
}
