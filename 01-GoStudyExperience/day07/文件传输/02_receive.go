package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main02() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//阻塞 等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("lintener.Accept err =", err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	fileName := string(buf[:n])

	//恢复"ok"
	conn.Write([]byte("ok"))

	//接收文件内容
	ReceFile(fileName, conn)

}

//ReceFile 接收用户发送来的文件
func ReceFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Creat err = ", err)
		return
	}

	buf := make([]byte, 1024*4)

	//接收多少，写多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err)
			}
			return
		}
		f.Write(buf[:n]) //往文件写内容
	}
}
