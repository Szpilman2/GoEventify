package main

import (
	"fmt"
	"goeventify/db"
	"goeventify/routes"
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

	db.InitDB()

	server := gin.Default()  // default returns an http engine instance with the logger and recovery middleware already attached.
	routes.RegisterRoutes(server)
	server.Run(":8080") //start listening to incoming requests to localhost:8080
}
