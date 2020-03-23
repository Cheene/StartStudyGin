package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
)

//静态文件
// html 页面上用到的样式文件，包括 css. js 和图片

//模板继承函数
func laodTemplates(templatesDir string) multitemplate.Renderer {
	//1 声明一个 NewRenderer
	r := multitemplate.NewRenderer()
	//2 加载 layout -- base.tmpl
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	//3 加载 includes
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	//为 layouts/ 和 includes/ 目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)

	}
	return r

}

func main() {
	r := gin.Default() // 创造一个默认路由引擎
	//模板的继承，借助 github.com/gin-contrib/multitemplate
	r.HTMLRender = laodTemplates("./templates/multitemplate")

	//加载静态文件
	r.Static("/mystatic", "./static")
	//设置自定义的模板函数
	// 一定是在解析模板文件之前
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//同时加载多个HTML模板
	//	r.LoadHTMLGlob("templates/**/*")
	r.GET("/index", func(c *gin.Context) {
		//定义模板
		//解析模板 这里有 . 代表着当前的目录
		r.LoadHTMLFiles("./templates/index.tmpl")
		//渲染模板 名字是模板的名字
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "none of your business",
		})
	})

	r.GET("posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "post none of your business",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title":       "get none of your business",
			"unsafe_code": "<alert> WARNING </alert>",
			"safe_code":   "<a href='https://www.nuoning.io'>踩踩踩</a>",
		})
	})

	r.GET("/mulindex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Hi": "Hi Index",
		})
	})

	r.GET("/mulhome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Hi": "Hi Home",
		})
	})

	r.Run(":8080")
}
