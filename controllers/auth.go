package controllers

import (
	"errors"
	"midterm/db"
	"midterm/env"
	"midterm/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	// Authenticate the user. Replace this with your actual authentication logic.
	user, err := authenticateAdmin(email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Create the JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	claims["user_id"] = user.UserId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with our secret
	tokenString, _ := token.SignedString([]byte(env.JwtSecret))

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authenticateAdmin(email, password string) (models.User, error) {
	user := models.User{}

	result := db.DbConnect.Where("email = ?", email).First(&user)
	if result.Error != nil {
		// fmt.Printf("Hello")
		return user, result.Error
	}

	if user.RoleId != 1 {
		return user, errors.New("unauthorized: user does not have the required role")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetAuthUser(c *gin.Context) (models.User, error) {
	var user models.User
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user found"})
		return user, errors.New("No user found")
	}

	if err := db.DbConnect.First(&user, int(id.(float64))).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return user, err
	}

	return user, nil
}