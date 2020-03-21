package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//解析模板
		t, err := template.ParseFiles("./hello.tmpl")
		if err != nil {
			fmt.Print("Err: %#v", err)
			return
		}
		user := Person{
			Name: "Chen",
			Age:  20,
		}
		m1 := map[string]interface{}{
			"name": "诺宁",
			"age":  18,
		}
		hobbyList := []string{"篮球", "足球", "乒乓球"}
		//渲染模板
		//传递多个变量的时候，通过 map 集合在一起，
		err = t.Execute(w, map[string]interface{}{
			"u1":        user,
			"m1":        m1,
			"hobbylist": hobbyList,
		})
		if err != nil {
			fmt.Printf("Error %#v", err)
		}
	})

	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		fmt.Print("Err: %#v", err)
		return
	}
}
