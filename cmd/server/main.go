package main

import (
	"log"
	"net/http"

	"github.com/Derikklok/go-practice-1/internal/db"
	"github.com/Derikklok/go-practice-1/internal/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// _ is here will be used becuase the package will be indirectly consumed by gorm
func main() {
	db.InitDB()
	// Set up Gin router
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Go Students Service running...",
		})
	})

	// routes
	r.GET("/students", handlers.GetAllStudents)
	r.POST("/students", handlers.CreateStudent)
	r.PUT("/students/:id", handlers.UpdateStudent)
	r.DELETE("/students/:id", handlers.DeleteStudent)

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
