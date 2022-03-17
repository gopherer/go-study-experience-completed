package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "通过自己创建的多路复用处理请求")
	if err != nil {
		fmt.Println(err)
	}
}

func main03() {
	mux := http.NewServeMux()
	mux.HandleFunc("/myMux", Handler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
