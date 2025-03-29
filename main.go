package main

import (
	"fmt"
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
	server.GET("/events", func (context *gin.Context)  {
		context.JSON(http.StatusOK, gin.H{"message":"Hello World!"})   //gin.H is a map (dictionary)
	})
	server.Run(":8080") //start listening to incoming requests to localhost:8080
}