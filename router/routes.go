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

	v1 := r.Group("/v1")
	{
		/// Todo ///
		v1.POST("/todo", todo.Create)
		v1.PUT("/todo/:id", todo.Update)
		v1.DELETE("/todo/:id", todo.Delete)
		v1.GET("/todo", todo.GetAll)
		v1.GET("/todo/:id", todo.GetOne)
		v1.GET("/todo/list", todo.List)
		// End Todo ///
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/todo/list", todo.List)
	}

}
