package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

//实现Handler接口
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "test...")
	if err != nil {
		fmt.Println(err)
	}
}
func main02() {
	myHandler := MyHandler{}
	//调用处理器
	http.Handle("/", &myHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
func main0201() {
	myHandler := MyHandler{}
	//调用处理器
	http.Handle("/hello", &myHandler)
	//创建Server结构，并详细配置里面的字段
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
