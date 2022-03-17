package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "你发送的请求地址是:", r.URL.Path)
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}
	_, err = fmt.Fprintln(w, "你发送的请求地址后查询字符串是:", r.URL.RawQuery)
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}
	_, err = fmt.Fprintln(w, "请求头中所有的信息:", r.Header)
	_, err = fmt.Fprintln(w, "请求头中Accept-Encoding的信息:", r.Header["Accept-Encoding"])
	_, err = fmt.Fprintln(w, "请求头中Accept-Encoding的属性值:", r.Header.Get("Accept-Encoding"))

	////获取请求体中的内容的长度
	//len := r.ContentLength
	////创建byte切片
	//body := make([]byte,len)
	////将请求体中的内容读到body中
	//r.Body.Read(body)
	////在浏览器中显示请求体中的内容
	//fmt.Fprintln(w,"请求体中的内容",string(body))
}
func main() {
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}
}
