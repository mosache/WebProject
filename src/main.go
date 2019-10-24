package main

import (
	api "WebProject/src/api/account"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/login", api.Login)

	app.Run()
}
