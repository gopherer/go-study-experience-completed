package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "hello world", r.URL.Path)
	if err != nil {
		fmt.Println(err)
	}
}
func main01() {
	http.HandleFunc("/", handler)
	//创建路由
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
