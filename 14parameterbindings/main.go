package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"u"`
	Password string `form:"password" json:"p"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		u := UserInfo{
			Username: username,
			Password: password,
		}

		fmt.Printf("%T : %v\n", u, u)
		fmt.Printf("%v : %v\n", u.Username, u.Password)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"value":   u,
		})
	})

	r.GET("/home", func(c *gin.Context) {
		var u UserInfo
		//需要传递地址，因为struct 是值类型

		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//因为不知道客户端会传递什么样的参数，所以需要反射
			fmt.Printf("%v \n", u)
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": u,
			})
		}
	})

	r.POST("/home", func(c *gin.Context) {
		var u UserInfo
		//需要传递地址，因为struct 是值类型
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//因为不知道客户端会传递什么样的参数，所以需要反射
			fmt.Printf("%v \n", u)
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": u,
			})
		}
	})
	r.Run(":8080")
}
