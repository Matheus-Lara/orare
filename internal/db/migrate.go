package db

import (
	"github.com/Matheus-Lara/orare/internal/api/logger"
	"github.com/Matheus-Lara/orare/internal/model"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	logger.Info("[Infrastructure] Migrating pending models...")
	db.AutoMigrate(
		&model.GoogleOAuthToken{},
		&model.User{},
	)
	logger.Info("[Infrastructure] Models migrated...")
}
