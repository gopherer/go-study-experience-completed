package main

import (
	"fmt"
	"net"
)

func main01() {
	//监听
	listener, err := net.Listen("tcp", ":8000") //默认为本机地址127.0.0.1
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err = ", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4*1024)

	n, err := conn.Read(buf)

	if err != nil || n == 0 {
		fmt.Println("conn.Read err = ", err)
		return
	}
	fmt.Printf("*%v*", string(buf[:n]))

}
