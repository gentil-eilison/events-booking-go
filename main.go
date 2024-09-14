package main

import (
	"github.com/gentil-eilison/events-booking-go/db"
	"github.com/gentil-eilison/events-booking-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Logger - logs operations of the server
	// Recovery - makes sure only part of the server app crashses when panicking
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // Will run on localhost:8080
}