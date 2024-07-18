package main

import (
	"example.com/Course_Project/db"
	"example.com/Course_Project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
