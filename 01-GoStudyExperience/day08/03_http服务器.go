package main

import "net/http"

func main03() {
	// //注册处理函数，用户连接。自动调用指定的处理函数
	//http.HandlerFunc("/", HandConn)//too many arguments in conversion to http.HandlerFunc

	// //监听 绑定
	// http.ListenAndServe(":8000", nil)

	mux := http.NewServeMux()

	handler := http.HandlerFunc(HandConn)

	mux.Handle("/hello", handler)

	http.ListenAndServe(":8000", mux)
}

//HandConn  w,给客户端回复的数据, reg 读取客户端发送的数据
func HandConn(w http.ResponseWriter, reg *http.Request) {
	w.Write([]byte("hello go")) //给客户端恢复数据
}
