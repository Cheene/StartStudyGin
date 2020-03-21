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
	t := template.New("f.tmpl").
		Delims("{<", ">}").
		//解析模板之前，要向模板对象中注册函数
		Funcs(template.FuncMap{
			"kua": kua,
		})
	//解析模板
	_, err := t.ParseFiles("./f.tmpl")

	if err != nil {
		fmt.Printf("Template Error : %#v", err)
		return
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
		return
	}
	//渲染模板
	t.Execute(w, "chenene")
}

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed,err: %#v\n", err)
		return
	}
	//渲染模板
	text := "This is index page"
	t.Execute(w, text)
}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("base_index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed,err: %#v\n", err)
		return
	}
	//渲染模板
	text := "This is index page"
	t.Execute(w, text)
}

func baseIndex(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tmpl, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		fmt.Printf("create template failed,err:%#0v\n", err)
		return
	}
	name := "彻恩"
	err = tmpl.ExecuteTemplate(w, "base_index.tmpl", name)
	if err != nil {
		fmt.Printf("render template failed,err: %#v\n", err)
		return
	}
	//渲染模板
}

func baseHome(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tmpl, err := template.ParseGlob("./templates/*.tmpl")
	if err != nil {
		fmt.Printf("create template failed,err:%#0v\n", err)
		return
	}
	name := "chene"
	err = tmpl.ExecuteTemplate(w, "base_home.tmpl", name)
	if err != nil {
		fmt.Printf("render template failed,err: \n", err)
		return
	}
	//渲染模板
}

func xss(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板

	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		// 参数 字符串，返回 template.HTML的格式
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("Err:%#v\n", err)
		return
	}
	jsStr := `<script>alert('黑呵呵')</script>`
	//渲染模板
	err = tmpl.Execute(w, jsStr)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/temp", demo1)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/baseIndex", baseIndex)
	http.HandleFunc("/baseHome", baseHome)
	http.HandleFunc("/xss", xss)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Err: %#v\n", err)
	}
}
