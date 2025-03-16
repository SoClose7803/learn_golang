package main

import (
	"user-api/config"
	"user-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.UserRoutes(r)

	r.Run(":8080") // Chạy server tại cổng 8080
	
}
