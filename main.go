package main

import (
	"fmt"
	"log"
	"midterm/db"
	"midterm/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	_, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		fmt.Printf("Connected to database\n")
	}
}

func main() {

	route := gin.Default()

	routers.InitRoute(route)

	route.Run(":8080")
}
