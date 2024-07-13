package db

import (
	"github.com/patrickishaf/githubservice/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitializeDb() {
	database, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	db = database
	dbError := db.AutoMigrate(&models.Repository{}, &models.Commit{})

	if dbError != nil {
		log.Println("failed to migrate database", dbError)
	}
}
