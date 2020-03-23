package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("user", "诺宁")
		addr := c.Query("addr")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  addr,
			"message":  "ok",
		})
	})
	r.Run(":8080")
}
