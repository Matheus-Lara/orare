package db

import (
	"github.com/Matheus-Lara/orare/internal/api/logger"
	"github.com/Matheus-Lara/orare/internal/model"
	"github.com/Matheus-Lara/orare/pkg/common"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB) {
	adminUser := model.User{
		Name:     "Orare",
		Email:    common.GetEnv("ADMIN_USER_EMAIL"),
		Password: common.GetEnv("ADMIN_USER_PASSWORD"),
		UserType: "ADMIN",
	}

	var existingAdminUser model.User
	db.Where("email = ?", adminUser.Email).First(&existingAdminUser)

	if existingAdminUser.ID != 0 {
		logger.Info("[Seed] Admin user already exists")
		return
	}

	result := db.Create(&adminUser)

	if result.Error != nil {
		logger.Fatal(result.Error.Error())
	}

	logger.Info("[Seed] Admin user created")
}
