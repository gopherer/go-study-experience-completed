package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//连接上服务器
	conn, err := net.Dial("tcp", "127.0.0.1:40000")
	if err != nil {
		panic(err)
	}
	//os.Stdin 标准输入 从键盘输入
	reader := bufio.NewReader(os.Stdin)
	b, err := reader.ReadSlice('\n')
	if err != nil {
		fmt.Println("终端输入失败", err)
		return
	}
	//conn.Write 返回一个写入的字节数和出现的错误
	n, err := conn.Write(b)
	if err != nil {
		fmt.Println("conn.Write err ", err)
		return
	}
	fmt.Printf("已将像服务器发送了%v个字节\n", n)
}
