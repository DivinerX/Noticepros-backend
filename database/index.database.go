package database

import (
	"fmt"
	"noticepros/config/db_config"
	"noticepros/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var errConnection error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_config.DB_HOST, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME, db_config.DB_PORT)

	DB, errConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})

	if errConnection != nil {
		panic("Failed connect to database.")
	}

	fmt.Println("Connected to database.")
}
