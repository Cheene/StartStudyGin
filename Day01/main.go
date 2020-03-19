package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")
	//不转换是字节 [60 104 49 62 72 101 108 108 111 32 87 111 114 108 100 33 33 33 60 47 104 49 62]
	//_, _ = fmt.Fprintln(w,b) //向谁写，写什么
	_, _ = fmt.Fprintln(w, string(b)) //向谁写，写什么
}

func main() {
	http.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Println("Http Server Failed,err : ", err)
		return
	}
}
