package db

import (
	"auth/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=auth_db user=postgres password=1234 dbname=auth port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect")
	}

	db.AutoMigrate(model.User{})

	DB = db
}
