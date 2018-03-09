package controller

import "github.com/gin-gonic/gin"

//HandleHomePage 首页控制器
func HandleHomePage(c *gin.Context) {
	data := getFrontDefault()
	data["PageTitle"] = "HomePage"
	c.HTML(200, "page/home", data)
}
