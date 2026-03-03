package db

import (
	"log"

	"github.com/Derikklok/go-practice-1/internal/models"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "./db/students.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto migrate tables
	DB.AutoMigrate(&models.Student{})
}
