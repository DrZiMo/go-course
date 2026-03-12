package routes

import "github.com/gin-gonic/gin"

func RegisterServer(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)
}
