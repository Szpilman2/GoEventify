package middlewars

import (
	"goeventify/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"}) //instead of json in middleware use: AbortWithStatusJSON
		return
	}

	user_id, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Not authorized"})
		return
	}

	context.Set("userId", user_id) //allows you to set some data in context to use it in next request handlers (in our case we can set userId)
	context.Next()  //make sure next request handler calls successfully.
}