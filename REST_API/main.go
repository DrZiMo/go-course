package main

import (
	"rest_api/db"
	"rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterServer(server)

	server.Run(":8080") // localhost:8080
}
