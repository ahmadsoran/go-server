package auth

import (
	"firstApp/helper"
	"firstApp/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var user models.User
	key := os.Getenv("SECRET_KEY")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	isVerified := helper.ComparePasswords(&user, []byte(user.Password))
	if !isVerified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message": user.Username + " logged in",
	})

}
