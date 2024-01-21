package controllers

import (
	"fmt"
	"log"
	"net/http"

	"midterm/db"     // replace "yourpackage" with the actual package name
	"midterm/models" // replace "yourpackage" with the actual package name

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DbC *gorm.DB

func init() {
	var err error
	DbC, err = db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		fmt.Printf("Connected to database\n")
	}
}

func CreateRole(c *gin.Context) {
	name := c.PostForm("name")

	role := models.Role{
		Name: name,
	}

	result := DbC.Create(&role)
	if result.Error != nil {
		// handle error, e.g. log it or return it in the HTTP response
		fmt.Println(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
