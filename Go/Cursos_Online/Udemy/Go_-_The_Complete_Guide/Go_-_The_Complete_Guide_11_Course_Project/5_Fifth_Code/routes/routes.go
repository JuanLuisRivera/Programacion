package routes

import (
	"example.com/Course_Project/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event", createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.POST("/event/:id/register", registerForEvent)
	authenticated.DELETE("/event/:id/register", cancelRegistration)

	server.DELETE("/event/:id", middlewares.Authenticate, deleteEvent)
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
