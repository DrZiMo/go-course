package routes

import "github.com/gin-gonic/gin"

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
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
