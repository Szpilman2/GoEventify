package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent) //dynamic identifier -> /events/1, /events/2
	server.POST("/events", CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
}