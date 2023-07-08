package main

import (
	"firstApp/conf"
	"firstApp/migration"
	"firstApp/router"

	"github.com/gin-gonic/gin"
)

func init() {
	conf.EnvLoader()
	conf.ConnectDB()
}
func main() {
	r := gin.Default()
	migration.Migration()
	router.Routes(r)
	r.Run("localhost:6969")
}
