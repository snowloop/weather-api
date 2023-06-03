package controllers

import (
	"github.com/gin-gonic/gin"
)

func ServeHome(gin_context *gin.Context) {
	gin_context.Writer.Write([]byte("<h1>Welcome to the main page of the weather api!</h1><a href=\"https://www.linkedin.com/in/alexandrejosse/\" target=\"_blank\">My Linkedin</a>"))
	gin_context.Writer.Header().Set("Content-Type", "text/html")
}
