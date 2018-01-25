package main

import (
	"controller"
	"os"
	"utils"
	"views"

	lg "github.com/imshuai/lightlog"

	"github.com/gin-gonic/gin"
)

func main() {
	lg.SetConsole(false)
	lg.SetLevel(lg.DEBUG)
	lg.SetPrefix("GoTyping")
	lg.SetRollingDaily("./logs", "GoTyping.log")
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery(), utils.Logger(os.Stdout))

	e.StaticFile("/favicon.png", "./statics/favicon.png")
	e.Static("/static", "./statics")
	e.SetFuncMap(views.TplFuncs)
	e.LoadHTMLGlob("./templates/**/*")
	e.GET("/", controller.HandleHomePage)
	e.GET("/page/:pid", controller.HandleArticalPagination)
	e.GET("/artical/:slug", controller.HandleArtical)

	manage := e.Group("/manage")
	{
		manage.GET("/", controller.HandleManageHomePage)
		manage.GET("/posts", controller.HandleManagePosts)
		manage.GET("/edit-post", controller.HandleManageEditPost)
		manage.GET("/create-post", controller.HandleManageCreatePost)
	}

	e.Run(":8080")

}
