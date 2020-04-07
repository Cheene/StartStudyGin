package main

import (
	"StartStudyGin/22bubble/controller"
	"StartStudyGin/22bubble/dao"
	"StartStudyGin/22bubble/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/**
1 创建引擎
2 连接数据库
3 增删改查
*/

func main() {
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		fmt.Printf("error: %v", err)
		panic(err)
	}
	defer dao.Close()
	//与数据表建立关联
	dao.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()
	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	vGroup := r.Group("v1")
	{
		//待办事项
		//增
		vGroup.POST("/todo", controller.UpdateAToDo)
		//删
		//删除某一个
		vGroup.DELETE("/todo/:id", controller.DeleteATodo)

		//改
		vGroup.PUT("/todo/:id", controller.ChangeATodo)

		//查
		//返回全部
		vGroup.GET("/todo", controller.GetATodo)
	}

	r.Run(":7070")
}
