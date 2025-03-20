package controller

import (
	"net/http"

	"github.com/Matheus-Lara/orare/internal/api"
	"github.com/Matheus-Lara/orare/internal/api/dto"
	"github.com/Matheus-Lara/orare/internal/api/service"
	"github.com/gin-gonic/gin"
)

type UserAuthController struct {
	service *service.UserAuthService
}

func (controller *UserAuthController) Register(c *gin.Context) {
	var createUserRequestDTO *dto.CreateUserRequestDTO
	api.ParseRequest(c, &createUserRequestDTO)

	user, err := controller.service.Register(createUserRequestDTO)

	if err != nil {
		api.ResponseBadRequest(c, err)
		return
	}

	api.ResponseSuccess(c, http.StatusCreated, user)
}

func (controller *UserAuthController) Login(c *gin.Context) {
	var loginRequestDTO *dto.LoginRequestDTO
	api.ParseRequest(c, &loginRequestDTO)

	token, err := controller.service.Login(loginRequestDTO)

	if err != nil {
		api.ResponseUnauthorized(c, err)
		return
	}

	api.ResponseSuccess(c, http.StatusOK, token)
}

func NewUserAuthController(service *service.UserAuthService) *UserAuthController {
	return &UserAuthController{
		service: service,
	}
}
