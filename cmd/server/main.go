package main

import (
	"net/http"

	"github.com/Derikklok/go-practice-1/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// _ is here will be used becuase the package will be indirectly consumed by gorm

var students []models.Student
var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "./db/students.db")
	if err != nil {
		panic("failed to connect to database")
	}
	// Migrate the schema (creates the table if it doesn't exist)
	db.AutoMigrate(&models.Student{})
}

func main() {
	InitDB()
	// Set up Gin router
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/students", func(ctx *gin.Context) {
		if err := db.Find(&students).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, students)
	})

	// Create a new student
	r.POST("/students", func(ctx *gin.Context) {
		var newStudent models.Student
		if err := ctx.ShouldBindJSON(&newStudent); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		// students = append(students, newStudent)
		// Save the students in db
		if result := db.Create(&newStudent); result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, newStudent)
	})

	r.Run(":8080")
}
