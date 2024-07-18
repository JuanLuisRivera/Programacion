package main

import (
	"net/http" // We add the http package to deliver web responses

	"github.com/gin-gonic/gin" // We add the gin package to achieve a basic web server functionality
)

func main() {
	server := gin.Default() // We use the gin package to create a basic "http" server structure

	server.GET("/event", getEvents) // We stablish a handler to execute with incoming requests

	server.Run(":8080") // We specify the server to listen to the port ":8080"
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello"}) // We use return the "ok" status via http (200), meaning everything worked correctly and a "map" struct
}
