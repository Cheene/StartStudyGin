package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取 path 参数
func main() {
	r := gin.Default()
	r.GET("/search/:user/:password", func(c *gin.Context) {
		username := c.Param("user")
		password := c.Param("password")

		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})

	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"year":    year,
			"month":   month,
		})
	})

	r.Run(":8080")
}
