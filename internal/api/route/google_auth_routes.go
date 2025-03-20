package route

import (
	"github.com/Matheus-Lara/orare/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func GoogleAuthRoutes(controller *controller.GoogleAuthController, group *gin.RouterGroup) {
	group.GET("/setup", controller.StartGoogleAuthorization)
	group.GET("/code", controller.AuthorizationCode)
}
