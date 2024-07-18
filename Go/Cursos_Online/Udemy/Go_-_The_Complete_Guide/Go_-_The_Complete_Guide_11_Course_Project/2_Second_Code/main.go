package main

import (
	"net/http"

	"example.com/Course_Project/db"
	"example.com/Course_Project/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/event", getEvents)
	server.POST("/event", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event // We populate the request with incoming data
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
	}

	event.Id = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created.", "event": event})
}
