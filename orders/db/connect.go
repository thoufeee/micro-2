package db

import (
	"log"
	"orders/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=order_db user=postgres password=1234 dbname=orders port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect")
	}

	db.AutoMigrate(model.Order{})

	DB = db
}
