package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/search", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	r.POST("/", func(c *gin.Context) {
		//user := c.PostForm("user")
		//addr := c.PostForm("addr")
		user := c.DefaultQuery("user", "诺宁")
		addr := c.DefaultQuery("addr", "***")
		c.JSON(http.StatusOK, gin.H{
			"name":    user,
			"address": addr,
		})
	})

	r.Run(":8080")
}
