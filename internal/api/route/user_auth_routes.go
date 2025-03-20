package route

import (
	"github.com/Matheus-Lara/orare/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func UserAuthRoutes(controller *controller.UserAuthController, group *gin.RouterGroup) {
	group.POST("/register", controller.Register)
	group.POST("/login", controller.Login)
}
