package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	e.StaticFile("/favicon.png", "./Statics/favicon.png")
	e.Static("/static", "./Statics")
	e.LoadHTMLGlob("./Templates/**/*")

	e.GET("/", routeIndex)
	e.GET("/page/:pid", routeIndex)

	e.Run(":8080")

}
