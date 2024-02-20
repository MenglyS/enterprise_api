package middleware

import (
	"fmt"
	"midterm/env"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Check for "Bearer" token type
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Initialize a new instance of `Claims`
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.JwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		// fmt.Printf("Claims: %v, type: %T\n", token.Claims, token.Claims)

		claimsMap, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		if int((*claimsMap)["role_id"].(float64)) > role {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		fmt.Printf("test", (*claimsMap)["user_id"])

		c.Set("id", (*claimsMap)["user_id"])

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
