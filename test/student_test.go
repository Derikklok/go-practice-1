package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Derikklok/go-practice-1/internal/db"
	"github.com/Derikklok/go-practice-1/internal/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("..")
	os.Mkdir("db", 0755)
	db.InitDB()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/students", handlers.GetAllStudents)
	r.POST("/students", handlers.CreateStudent)
	r.PUT("/students/:id", handlers.UpdateStudent)
	r.DELETE("/students/:id", handlers.DeleteStudent)
	return r
}

func TestCreateStudent(t *testing.T) {
	r := SetupRouter()

	// Creating a student in JSON format
	var json = []byte(`{"id": "1", "name": "John Doe", "grade": "A"}`)
	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllStudents(t *testing.T) {
	r := SetupRouter()

	req, _ := http.NewRequest("GET", "/students", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
