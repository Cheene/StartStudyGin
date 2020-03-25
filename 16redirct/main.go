package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	//HTTP重定向
	r.GET("/home", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://cheene.github.io/")
	})

	//路由重定向
	r.GET("/login", func(c *gin.Context) {
		//处理哪一个呢？
		c.Request.URL.Path = "/index"
		r.HandleContext(c) // 继续处理这个 context
	})
	r.Run(":8080")
}
