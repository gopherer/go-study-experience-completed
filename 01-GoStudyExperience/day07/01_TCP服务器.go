package main

import (
	"fmt"
	"net"
)

func main01() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer listener.Close()

	//阻塞等待用户链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//接收用户请求
	buf := make([]byte, 1024) //1024大小的缓冲区
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf = ", string(buf[:n]))

	defer conn.Close() //关闭当前用户链接

}

//当有新的变量生成的时候，如果有重复的变量那么那个重复的变量就会由定义变为赋值。
//变量的声明必须有新的变量生成否则不能通过编译器。
