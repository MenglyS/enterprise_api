package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DbConnect *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables")
	}

	host := os.Getenv("HOSTNAME")
	user := os.Getenv("NAME")
	dbname := os.Getenv("DB")
	password := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, port)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	DbConnect = db

	return DbConnect, nil
}
