package routes

import (
	"rest_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterServer(server *gin.Engine) {
	// User
	server.GET("/users", getUsers)
	server.GET("/users/:id", getSingleUser)
	server.POST("/users", createUser)
	server.POST("/users/login", loginUser)
	server.PUT("/users/:id", updateUser)
	server.DELETE("/users/:id", deleteUser)

	// Event
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", middleware.Authenticate, createEvent)
	server.PUT("/events/:id", middleware.Authenticate, updateEvent)
	server.DELETE("/events/:id", middleware.Authenticate, deleteEvent)
}
