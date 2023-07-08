package router

import (
	"firstApp/router/auth"
	"firstApp/router/todo"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	/// Auth
	r.POST("/login", auth.Login)
	r.POST("/signup", auth.Signup)
	/// End Auth ///
	// auth middleware
	r.Use(auth.AuthMiddleware())
	/// Todo ///
	r.POST("/todo", todo.Create)
	r.PUT("/todo/:id", todo.Update)
	r.DELETE("/todo/:id", todo.Delete)
	r.GET("/todo", todo.GetAll)
	r.GET("/todo/:id", todo.GetOne)
	// End Todo ///

}
