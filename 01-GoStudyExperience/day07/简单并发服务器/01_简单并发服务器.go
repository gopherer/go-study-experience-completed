package main

import (
	"fmt"
	"net"
	"strings"
)

func main01() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		//处理用户请求
		go HandleConn(conn)
	}
}

//HandleConn 处理用户请求
func HandleConn(conn net.Conn) {
	//函数调用完毕，关闭函数
	defer conn.Close()

	//获取客户端网络地址
	addr := conn.RemoteAddr().String()//转换成字符串
	fmt.Println(addr, "connect sucessful")

	buf := make([]byte, 2048)
	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s] : %s\n", addr, string(buf[:n]))

		//if "exit" == string(buf[:n-1]) {//使用与necat -1为了消除\n
		if "exit" == string(buf[:n-2]) { //-2 是消除Windows文本下\r\n
			fmt.Println(addr, "exit")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))//小写字母 转变为 大写字母发送给客户端
	}
}