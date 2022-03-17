package main

import (
	"fmt"
	"net"
	"os"
)

func main02() {
	//主动连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	//mian函数调用完毕，关闭conn连接
	defer conn.Close()

	//匿名函数---- 给服务器发送内容
	go func() {
		//储存内容的空间
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf) //从键盘读取内容，保存在buf中 输入exit退出
			if err != nil {
				fmt.Println("os.Stdin err = ", err)
				return
			}
			//给服务器发送内容
			conn.Write(buf[:n])
		}
	}()

	//接收服务器发送回来的内容
	//切片缓存
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收服务器数据
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}