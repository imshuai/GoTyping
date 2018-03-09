package controller

import (
	"github.com/gin-gonic/gin"
)

func getFrontDefault() gin.H {
	data := gin.H{}
	data["SiteName"] = "GoTyping"
	return data
}

func getBackDefault() gin.H {
	return gin.H{}
}
