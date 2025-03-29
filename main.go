package main

import (
	"fmt"
	"goeventify/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	asciArt := `
	██████╗  ██████╗ ███████╗██╗   ██╗███████╗███╗   ██╗████████╗██╗███████╗██╗   ██╗
	██╔════╝ ██╔═══██╗██╔════╝██║   ██║██╔════╝████╗  ██║╚══██╔══╝██║██╔════╝╚██╗ ██╔╝
	██║  ███╗██║   ██║█████╗  ██║   ██║█████╗  ██╔██╗ ██║   ██║   ██║█████╗   ╚████╔╝ 
	██║   ██║██║   ██║██╔══╝  ╚██╗ ██╔╝██╔══╝  ██║╚██╗██║   ██║   ██║██╔══╝    ╚██╔╝  
	╚██████╔╝╚██████╔╝███████╗ ╚████╔╝ ███████╗██║ ╚████║   ██║   ██║██║        ██║   
	╚═════╝  ╚═════╝ ╚══════╝  ╚═══╝  ╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝╚═╝        ╚═╝   
	`
	fmt.Println(asciArt)

	server := gin.Default()  // default returns an http engine instance with the logger and recovery middleware already attached.
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") //start listening to incoming requests to localhost:8080
}

func getEvents(context *gin.Context)  {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)   //gin.H is a map (dictionary)
}

func createEvent(context *gin.Context)  {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	context.JSON(http.StatusOK, gin.H{"message": "Event Created!", "event": event})
}