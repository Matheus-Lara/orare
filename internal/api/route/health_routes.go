package route

import (
	"github.com/Matheus-Lara/orare/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func HealthRoutes(controller *controller.HealthController, group *gin.RouterGroup) {
	group.GET("/health", controller.GetHealth)
}
