package main

import (
	"fmt"
	"net/http"
)

func main04() {
	mux := http.NewServeMux()

	handle := http.HandlerFunc(HandleConn)

	mux.Handle("/go", handle)

	http.ListenAndServe(":8000", mux)
}

//HandleConn w 服务器给客户端回复的数据 reg 客户端个服务器发送的数据
func HandleConn(w http.ResponseWriter, reg *http.Request) {
	w.Write([]byte("hello go"))
	fmt.Println("reg.Method = ", reg.Method)
	fmt.Println("reg.URL = ", reg.URL)
	fmt.Println("reg.Header = ", reg.Header)
	fmt.Println("reg.Body = ", reg.Body)
}
