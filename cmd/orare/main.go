package main

import (
	"log"

	"github.com/Matheus-Lara/orare/internal/db"
	"github.com/Matheus-Lara/orare/internal/i18n"
	"github.com/Matheus-Lara/orare/internal/server/di"
	"github.com/Matheus-Lara/orare/pkg/common"
	"github.com/Matheus-Lara/orare/pkg/environment"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("[Orare] Service Started")

	godotenv.Load()
	i18n.Init()

	//setupDatabase()
	startHttpServer()

	common.WaitOsInterruption()
}

func startHttpServer() {
	httpServer := di.InitializeHttpServer()

	log.Println("[HttpServer] Starting...")
	go httpServer.Run()
	log.Println("[HttpServer] Started")
}

func setupDatabase() {
	log.Println("[Infrastructure] Connecting to database...")
	gormDbConnection := db.Init()
	if environment.IsDevelopment() {
		db.MigrateModels(gormDbConnection)
		db.SeedAdminUser(gormDbConnection)
	}
	log.Println("[Infrastructure] Database connected...")
}
