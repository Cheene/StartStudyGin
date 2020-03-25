package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创造默认路由
	r := gin.Default()
	//加载资源
	r.LoadHTMLFiles("./404.html")
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	//处理所有的请求
	r.Any("/home", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"method": "POST",
			})

		}
	})
	//访问不存在的请求
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//；路由组的使用
	classGroup := r.Group("/class")
	{
		classGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"methdd": "/class/index",
			})
		})
		classGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"methdd": "/class/login",
			})
		})
		classGroup.GET("/pre", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"methdd": "/class/pre",
			})
		})
		classGroup.POST("/pre", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"methdd": "/class/pre_POST",
			})
		})
	}

	//路由的嵌套
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/shop/index",
			})
		})

		xx := shopGroup.Group("/sheet")
		xx.GET("/male", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/shop/sheet/male",
			})
		})
		xx.GET("/woman", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/shop/sheet/woman",
			})
		})
	}
	r.Run(":8080")
}
