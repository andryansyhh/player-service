package middleware

import (
	"net/http"
	"os"
	"player-service/src/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Change this to a secure secret key
// GenerateJWT generates a new JWT token
func GenerateJWT(userUuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_uuid"] = userUuid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TokenCheckerMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	splitToken := strings.Split(token, " ")

	// Check if the token exists and matches the secret token
	if token == "" || splitToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	tkn, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// getredis
	err = config.GetRedisData(splitToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok || !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	userUuid, ok := claims["user_uuid"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User UUID not found in token"})
		c.Abort()
		return
	}

	c.Set("user_uuid", userUuid)

	c.Next()
}
