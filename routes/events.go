package routes

import (
	"fmt"
	"goeventify/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //gives parameter as string
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events) //gin.H is a map (dictionary)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	// event.ID = 1
	// event.UserID = 1

	err = event.Save()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. Try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Created!", "event": event})
}

func UpdateEvent(context *gin.Context)  {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //gives parameter as string
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	var UpdatedEvent models.Event
	err = context.ShouldBindJSON(&UpdatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}
	UpdatedEvent.ID = eventId
	err = UpdatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func DeleteEvent(context *gin.Context)  {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //gives parameter as string
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}