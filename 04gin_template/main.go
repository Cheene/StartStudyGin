package main

import (
	"fmt"

	//	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 遇事不决，写注释
// 1 定义模板
// 2 解析模板
// 3 渲染模板
func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	temp, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Err for template: %#v", err)
		return
	}
	//渲染模板
	err = temp.Execute(w, temp)
	if err != nil {
		fmt.Printf("Err for template: %#v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Start Failed.Err: %#v", err)
	}
	//r := gin.Default()
	//r.GET("gintemplate",getTemplate)
	//
	//r.Run(":8080")
}
