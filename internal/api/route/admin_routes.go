package route

import (
	"github.com/Matheus-Lara/orare/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(controller *controller.AdminController, adminGroup *gin.RouterGroup) {
	adminGroup.POST("/migrate", controller.MigrateModels)
}
