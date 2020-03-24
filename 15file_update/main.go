package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"time"
)

//gin 文件的提交

func main() {
	//1 定义路由
	r := gin.Default()
	//2 解析模板
	r.LoadHTMLFiles("./index.html", "./home.html")
	//3 响应请求
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	//单文件的上传
	r.POST("/sfileupload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println("The fileName is ", file.Filename)
		now := time.Now().Unix()
		//需要保证 D:/tmp/ 的目录是存在的
		dst := fmt.Sprint("D:/tmp/%s\n", string(now)+file.Filename)
		//上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"location": "D:/tmp/" + string(now) + file.Filename,
		})
	})

	//多文件的上传
	r.POST("/dfileupload", func(c *gin.Context) {
		//1 获取多个文件
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		//2 获取文件
		files := form.File["file"]
		//3 遍历每一个文件，并且保存
		for index, file := range files {
			log.Println(index, " : ", file.Filename)
			dst := path.Join("./", file.Filename)
			//开始上传
			c.SaveUploadedFile(file, dst)
		}
		//4 返回并显示上传成功的文件的个数
		c.JSON(http.StatusOK, gin.H{
			"message":     "ok",
			"file_number": len(files),
		})
	})
	//4 启动
	r.Run(":8080")

}
