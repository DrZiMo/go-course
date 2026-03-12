package main

import (
	"rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterServer(server)

	server.Run(":8080") // localhost:8080
}
