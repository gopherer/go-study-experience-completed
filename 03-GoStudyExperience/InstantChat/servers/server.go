package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听：")
	listen, err := net.Listen("tcp", "127.0.0.1:40000")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := listen.Close()
		if err != nil {
			panic(err)
		}
	}()
	for {
		fmt.Println("等待用户连接:")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("该用户连接失败.")
			continue
		}
		//处理用户连接
		go Process(conn)
	}
}
func Process(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("该用户关闭失败.")
		}
	}()
	fmt.Printf("%v该用户已连接上服务器\n", conn.RemoteAddr().String())
	for {
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1、等待客户端通过conn发送数据
		//2、如果客户端没有发送数据conn.Read(buf)就会阻塞
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器Read err ", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
