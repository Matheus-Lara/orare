package controller

import (
	"net/http"

	"github.com/Matheus-Lara/orare/internal/api"
	"github.com/Matheus-Lara/orare/internal/api/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func (controller *UserController) Profile(c *gin.Context) {
	user, err := controller.service.Profile(c)

	if err != nil {
		api.ResponseNotFound(c, err)
		return
	}

	api.ResponseSuccess(c, http.StatusCreated, user)
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}
