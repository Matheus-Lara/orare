package controller

import (
	"net/http"

	"github.com/Matheus-Lara/orare/internal/api"
	"github.com/Matheus-Lara/orare/internal/db"
	"github.com/Matheus-Lara/orare/internal/i18n"
	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (controller *AdminController) MigrateModels(c *gin.Context) {
	db.MigrateModels(db.GetConnection())
	api.ResponseSuccess(c, http.StatusCreated, gin.H{"message": i18n.Message("Admin.MigrateModels.Response.Success")})
}

func NewAdminController() *AdminController {
	return &AdminController{}
}
