package route

import (
	"github.com/Matheus-Lara/orare/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(controller *controller.UserController, group *gin.RouterGroup) {
	group.GET("/profile", controller.Profile)
}
