package controller

import (
	"net/http"

	"github.com/Matheus-Lara/orare/internal/api"
	"github.com/Matheus-Lara/orare/internal/api/service"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	service *service.HealthService
}

func (controller *HealthController) GetHealth(c *gin.Context) {
	api.ResponseSuccess(c, http.StatusOK, controller.service.GetHealth())
}

func NewHealthController(service *service.HealthService) *HealthController {
	return &HealthController{
		service: service,
	}
}
