//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Matheus-Lara/orare/internal/api/controller"
	"github.com/Matheus-Lara/orare/internal/api/service"
	"github.com/Matheus-Lara/orare/internal/db"
	"github.com/Matheus-Lara/orare/internal/repository"
	"github.com/Matheus-Lara/orare/internal/server"
	"github.com/google/wire"
)

func InitializeHttpServer() *server.HttpServer {
	wire.Build(
		server.NewHttpServer,
		controller.NewAdminController,
		controller.NewGoogleAuthController,
		controller.NewHealthController,
		controller.NewUserAuthController,
		controller.NewUserController,
		service.NewGoogleAuthService,
		service.NewHealthService,
		service.NewJWTService,
		service.NewUserAuthService,
		service.NewUserService,
		repository.NewUserRepository,
		db.GetConnection,
	)
	return &server.HttpServer{}
}
