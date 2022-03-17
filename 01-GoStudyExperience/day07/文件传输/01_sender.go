package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main01() {
	//提示用户输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)

	//获取文件名 info.Name()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	//主动连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	defer conn.Close()

	//给接收方，发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}

	//接收方收到信息，如果回复"ok",就说明对面准备好了可以发送文件
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	if "ok" == string(buf[:n]) {
		//可以开发送文件内容
		SendFile(path, conn)
	}

}

//SendFile 发送文件
func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)

	//读文件内容，都多少就发多少
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("os.Stat err = ", err)
			}
			return
		}
		//发送文件
		conn.Write(buf[:n]) //给服务器发送文件
	}

}
