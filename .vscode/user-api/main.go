package main

import (
	"fmt"

	"user-api/config"
	"user-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server is starting...") 

	config.ConnectDatabase()

	r := gin.Default()
	routes.UserRoutes(r)

	r.Run(":8080") // Chạy server tại cổng 8080
}
