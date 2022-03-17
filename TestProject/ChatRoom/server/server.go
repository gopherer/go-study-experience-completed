package main

import (
	"Project/ChatRoom/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:40000")
	if err != nil {
		fmt.Println("net listen err: ", err)
		return
	}
	//延时关闭
	defer func() {
		err = listen.Close()
		if err != nil {
			fmt.Println("listen close err:", err)
			os.Exit(1)
		}
	}()
	for {
		fmt.Println("server is listening: ")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			continue
		}
		//开启一个协程来处理用户连接
		go process(conn)
	}

}

//处理用户连接
func process(conn net.Conn) {
	mes := message.Message{}
	loginMes := message.LoginMes{}
	//定义消息管道
	//mesChan := make(chan message.Message)
	//定义错误处理管道
	errChan := make(chan error)
	defer func() {
		err := conn.Close()
		if err != nil {
			err = errors.New("conn close err")
			errChan <- err
			return
		}
	}()
	//处理校验信息
	verify := make([]byte, 4)
	n, err := conn.Read(verify)
	if n != 4 || err != nil {
		err = errors.New("verify conn read err")
		errChan <- err
	}
	fmt.Println(verify)
	//处理用户发送的mesInfo
	pkg := binary.BigEndian.Uint32(verify)
	userInfo := make([]byte, pkg)
	n, err = conn.Read(userInfo)
	if n != int(pkg) || err != nil {
		err = errors.New("userInfo conn read err")
		errChan <- err
	}
	//json反序列化接收到的userInfo
	err = json.Unmarshal(userInfo, &mes)
	if err != nil {
		err = errors.New("json unmarshal err")
		errChan <- err
	}
	//json反序列话mes.Data 得到loginMes
	//mesData := len(mes.Data)
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		err = errors.New("json unmarshal err")
		errChan <- err
	}
	fmt.Println(loginMes)
	return
}
