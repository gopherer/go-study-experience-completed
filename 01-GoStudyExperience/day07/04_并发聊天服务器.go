package main

import (
	"fmt"
	"net"
	"strings"
)

//Client 用户信息
type Client struct {
	C    chan string //用于发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

//保存在线用户  cliAddr======>Client
var onlineMap map[string]Client

var message = make(chan string)

func main04() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000") //默认地址为127.0.0.1
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	fmt.Println("Server have opened")

	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，就遍历map，给map每个成员都发送消息
	go Mannager()
	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}
		go HanledConn(conn) //处理用户连接
	}

}

//HanledConn 处理用户连接
func HanledConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体
	cli := Client{make(chan string), cliAddr, cliAddr} //结构体赋值 要使用{} 而不是()
	fmt.Println(cli.Addr, "User login")
	//把结构体添加到map
	onlineMap[cliAddr] = cli
	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)
	//广播某个人在线
	message <- MakeMsg(cli, "login")
	message <- "This message is a login prompt"
	isQuit := make(chan bool) //对方是否主动退出

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //当前用户已断开或出问题
				fmt.Println(MakeMsg(cli, "User has been offline"))
				message <- MakeMsg(cli, "User has been offline")
				isQuit <- true
				return
			}
			if err != nil {
				fmt.Println("conn.Read err = ", err)
				MakeMsg(cli, "User has been error")
				return
			}

			//msg := string(buf[:n-1]) //通过Windows，nc测试会多一个换行符
			msg := string(buf[:n-2]) //通过Windows,-2 是消除Windows文本下\r\n

			if len(msg) == 3 && msg == "who" {
				//遍历map,给当前用户发送所有成员
				conn.Write([]byte("user list: \n"))
				for _, tep := range onlineMap {
					msg = tep.Addr + ":" + tep.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//rename| 新的名字
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli //不可删除
				conn.Write([]byte("rename ok\n"))

			} else {
				//转发此消息
				message <- MakeMsg(cli, msg)
			}

		}
	}()
	for { //保证当前用户的连接
		//通过select检测channel的数据流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)
			return
		}
	}
}

//Mannager 转发消息，只要有消息来了，就遍历map，给map每个成员都发送消息
func Mannager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息，就会阻塞
		//就遍历map，给map每个成员都发送消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

//WriteMsgToClient 发送信息给客户
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

//MakeMsg 提示用户上线信息
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}
