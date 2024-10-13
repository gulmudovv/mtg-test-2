package database

import (
	"MTG-test-2/server/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "12345"
	dbname   = "ginblog"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//connect database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	DB = db
	log.Println("Connected to database!")

	err = db.AutoMigrate(&models.Message{})
	if err != nil {
		log.Fatalf("Migrate to database failed: %v", err)
	}

}

func Create(id, data string) {
	DB.Create(&models.Message{
		Socket: id,
		Data:   data,
	})
}
