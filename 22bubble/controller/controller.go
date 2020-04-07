package controller

import (
	"StartStudyGin/22bubble/dao"
	"StartStudyGin/22bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
url --> controller --> logic --> model
请求	-》 控制器 	->	业务逻辑 --》模型层的增删改查
 */
*/
func IndexHandler (c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}

func UpdateAToDo(c *gin.Context) {
	//请求中获取数据
	var todo models.Todo
	c.BindJSON(&todo)
	//存入数据
	if err := models.CreateATodo(&todo);err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else {
		//返回相应状态
		c.JSON(http.StatusOK,todo)
	}
}

func DeleteATodo(c *gin.Context) {
	//请求中获取数据
	id := c.Params.ByName("id")
	var todo Todo
	if err = dao.DB.Where("id=?",id).Error;err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error":"id invalid",
		})
		return
	}
	c.BindJSON(&todo)
	//这里不写 where 条件会全部删除
	if err = DB.Where("id=?",id).Delete(&todo).Error;err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			id:"deleted",
		})
	}
	//删除数据
	//返回相应状态
}

func ChangeATodo(c *gin.Context) {
	//请求中获取数据
	id := c.Params.ByName("id")
	//更新数据
	var todo Todo
	if err = DB.Where("id=?",id).First(&todo).Error; err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error":"id invalid",
		})
		return
	}
	c.ShouldBind(&todo)
	if err = DB.Save(&todo).Error; err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":"success",
		})
	}
	//返回相应状态
}

func GetATodo(c *gin.Context) {
	todoList,err := models.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,todoList)
	}
}
