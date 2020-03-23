package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//json
// c.json 是 json 的初始化
func main() {
	//1 定义默认引擎
	r := gin.Default()
	//2 函数
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hollo world",
		})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		//结构体直接返回
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "是你的诺宁"
		msg.Message = "Hey NUO NING"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	//3 跑起来
	r.Run(":8080")
}
