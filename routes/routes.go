package routes

import (
	"goeventify/middlewars"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//first approach to add middleware: add middleware as second argument like: server.POST("/events", middlewars.Authenticate ,CreateEvent)
	// server.GET("/events", GetEvents)
	// server.GET("/events/:id", GetEvent) //dynamic identifier -> /events/1, /events/2
	// server.POST("/events", middlewars.Authenticate ,CreateEvent)
	// server.PUT("/events/:id", UpdateEvent)
	// server.DELETE("/events/:id", DeleteEvent)

	//second way is to create group like below:
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent) //dynamic identifier -> /events/1, /events/2
	authenticated := server.Group("/")
	authenticated.Use(middlewars.Authenticate)
	authenticated.POST("/events" ,CreateEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	server.POST("/signup", SignUp)
	server.POST("/login", Login)
}