package db

import (
	"log"
	"products/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=products_db user=postgres password=1234 dbname=product port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect")
	}

	db.AutoMigrate(&model.Product{})

	DB = db
}
