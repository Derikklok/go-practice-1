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


