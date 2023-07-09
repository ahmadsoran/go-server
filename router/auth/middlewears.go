package auth

import (
	"firstApp/helper"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Do some stuff here
		// cookie token check
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{
				"message": "you are not logged in",
			})
			c.Abort()
			return
		}
		secretKey := os.Getenv("SECRET_KEY")
		// token check
		isVerified, err := helper.VerifyJWTToken(token, secretKey)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		if !isVerified.Valid {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		user := isVerified.Claims.(jwt.MapClaims)
		id := user["id"].(float64)
		c.Set("user_id", id)
		c.Next()
	}
}
