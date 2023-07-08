package router

import (
	"firstApp/router/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/login", auth.Login)
	r.POST("/signup", auth.Signup)
}
