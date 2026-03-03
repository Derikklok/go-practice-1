// internal/handlers/student.go
package handlers

import (
	"net/http"

	"github.com/Derikklok/go-practice-1/internal/db"
	"github.com/Derikklok/go-practice-1/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(ctx *gin.Context) {
	var students []models.Student
	if err := db.DB.Find(&students).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	// convert the header values first then the body
	ctx.JSON(http.StatusOK, students)
}

func CreateStudent(ctx *gin.Context) {
	var newStudent models.Student
	// use & when values should be actullay needs tobe copied into memory
	if err := ctx.ShouldBindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// Save the record to db
	if err := db.DB.Create(&newStudent).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newStudent)
}

// func CreateStudent(c *gin.Context) {
// 	var newStudent db.Student
// 	if err := c.ShouldBindJSON(&newStudent); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.DB.Create(&newStudent).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, newStudent)
// }
