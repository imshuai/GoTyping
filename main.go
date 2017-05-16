package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	e.GET("/", func(c *gin.Context) {
		c.String(200, "%v", "Hello World")
	})

	e.Run(":8080")

}
