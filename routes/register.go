package routes

import (
	"goeventify/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //gives parameter as string
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user", "err" : err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registered on event successfully"})
}

func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //gives parameter as string
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration", "err" : err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Canceled registration on event successfully"})
}