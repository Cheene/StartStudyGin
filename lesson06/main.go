package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数 kua,多返回值的时候，第二个返回值必须是Error类型
	kua := func(name string) (string, error) {
		return "夸一下" + name, nil
	}

	//定义模板
	//① 创建一个名字为 f 的模板对象
	//② 解析模板{注意模板对象的名字与解析的文件名字要对应上}
	t := template.New("f.tmpl")
	//解析模板之前，要向模板对象中注册函数
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	//解析模板
	_, err := t.ParseFiles("./f.tmpl")

	if err != nil {
		fmt.Printf("Template Error : %#v", err)
	}
	name := "诺宁"
	//渲染模板
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err: %#v\n", err)
	}
	//渲染模板
	t.Execute(w, "chenene")
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/temp", demo1)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Err: %#v\n", err)
	}

}
