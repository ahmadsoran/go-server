package auth

import (
	"firstApp/conf"
	"firstApp/helper"
	"firstApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed := helper.HashAndSal([]byte(user.Password))
	user.Password = hashed

	// Create a new user record in the database
	result := conf.DB.Create(&user)
	if result.Error != nil {
		if helper.IsSqlDuplicateError(result.Error) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
