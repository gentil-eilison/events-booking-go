package main

import (
	"net/http"

	"github.com/gentil-eilison/events-booking-go/db"
	"github.com/gentil-eilison/events-booking-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Logger - logs operations of the server
	// Recovery - makes sure only part of the server app crashses when panicking
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // Will run on localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = 1
	event.ID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}