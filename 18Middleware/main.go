package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//定义中间件 统计耗时
func m1(c *gin.Context) {
	start := time.Now()
	fmt.Println("m1 in...")
	//计时
	c.Next() // 调用后面的程序
	//c.Abort() //阻止后面的程序
	cost := time.Now().Sub(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out....")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "nuoning")
	/*
		1 m2 会继续执行
		2 不会继续执行接下来的函数。
	*/
	//c.Abort() // 阻止后面的事件
	fmt.Println("m2 out...")
}

//通常使用闭包来实现中间件的形式
func authMildWare(check bool) gin.HandlerFunc {
	// 连接数据库，或其他准备
	return func(c *gin.Context) {
		//判断是否是登录用户
		//如果是登录用户
		if check {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()

	//全局注册中间件函数
	r.Use(m1, m2)

	r.GET("/index", authMildWare(true), func(c *gin.Context) {
		fmt.Println("index")
		name, ok := c.Get("name") //跨中间件来取值
		if !ok {
			name = "annno"
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": name,
		})
	})

	r.GET("/shop", authMildWare(false), func(c *gin.Context) {
		fmt.Println("shop")
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		fmt.Println("user")
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	//路由组创建中间件
	xxGroup := r.Group("/xx", authMildWare(true)) //路由组中间件注册方法1
	xxGroup.Use(m1)                               //路由组中间件注册方法2
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "true_/x/index",
			})
		})
	}

	r.Run(":9090")
}
