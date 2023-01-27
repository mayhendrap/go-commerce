package migrations

import (
	"go-commerce/entities"
	"gorm.io/gorm"
	"log"
)

func MigrationModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Product{},
	)
	if err != nil {
		log.Fatalln("failed to migrate models!.")
	}
}
