package main

import (
	"fmt"
	"go-commerce/configs"
	"go-commerce/migrations"
	"go-commerce/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	_, db, err := configInitialization()
	if err != nil {
		log.Fatal("Config initialization failed!.")
	}

	migrations.MigrationModels(db)

	router := routes.SetupRouter()
	route := router.Group("/")

	routes.ProductRoute(db, route)

	err = router.Run() // wil run default in port 8080
	if err != nil {
		log.Fatal("Failed to run app: ", err)
	}

}

func configInitialization() (configs.Config, *gorm.DB, error) {
	var (
		config configs.Config
		db     *gorm.DB
		err    error
	)

	// Load configuration
	config, err = configs.LoadConfig(".")
	if err != nil {
		return config, nil, err
	}

	// Connect database
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort, config.DbTz,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return config, db, err
	}
	log.Println("[DB] - Connection successfully established")
	return config, db, err
}
