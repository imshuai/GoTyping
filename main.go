package main

import (
	"views"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	e.StaticFile("/favicon.png", "./Statics/favicon.png")
	e.Static("/static", "./Statics")
	e.SetFuncMap(views.TplFuncs)
	e.LoadHTMLGlob("./Templates/**/*")
	e.GET("/", HandleHomePage)
	e.GET("/page/:pid", HandleArticalPagination)
	e.GET("/artical/:slug", HandleArtical)

	manage := e.Group("/manage")
	{
		manage.GET("/", HandleManageHomePage)
		manage.GET("/posts", HandleManagePosts)
		manage.GET("/edit-post", HandleManageEditPost)
		manage.GET("/create-post", HandleManageCreatePost)
	}

	e.Run(":8080")

}
