package bootstrap

import (
	"log"

	//"github.com/gabrielfmcoelho/abare-api/domain"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
	//&domain.User{},
	// Add other models here
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}
}
