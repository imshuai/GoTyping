package main

import "github.com/gin-gonic/gin"
import "strconv"
import "net/http"

func routeIndex(c *gin.Context) {
	var err error
	pid := 1
	if len(c.Params) > 0 {
		pid, err = strconv.Atoi(c.Param("pid"))
		if err != nil {
			pid = 1
		}
	}
	nav := NavbarData{
		MainPage: func() bool {
			if pid == 1 {
				return true
			}
			return false
		}(),
	}
	nav.Categories, err = GetCategories(-1)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	header := HeaderData{
		PageTitle: func() string {
			if pid == 1 {
				return "首页 - GoTyping"
			}
			return "第" + strconv.Itoa(pid) + "页 - GoTyping"
		}(),
		CSS: []string{},
	}
	summary := GetSummariesWithPageID(pid)
	prepage := ""
	nextpage := "2"
	c.HTML(http.StatusOK, "Summary/index.html", gin.H{
		"head":   header,
		"navbar": nav,
		"script": []string{},
		"option": Option{},
		"data": gin.H{
			"summary":  summary,
			"prepage":  prepage,
			"nextpage": nextpage,
		},
	})
}
