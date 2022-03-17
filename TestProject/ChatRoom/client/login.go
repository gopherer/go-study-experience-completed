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

func login(userID int, userPwd string) (err error) {
	//接收用户
	fmt.Print("用户名:")
	_, err = fmt.Scanln(&userID)
	if err != nil {
		return errors.New("系统出错，请重新输入")
	}
	fmt.Print("用户密码:")
	_, err = fmt.Scanln(&userPwd)
	if err != nil {
		return errors.New("系统出错，请重新输入")
	}
	//指定协议
	conn, err := net.Dial("tcp", "127.0.0.1:40000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return errors.New("net Dial err")
	}
	//延时关闭用户连接
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("conn close err:", err)
			os.Exit(1)
		}
	}()
	//定义Message结构体 用于发用数据给server
	mes := message.Message{
		Type: message.LoginType, //确定发送数据类型
		Data: "",
	}
	//定义LoginMes结构体
	login := message.LoginMes{
		UserId:   userID,
		UserPwd:  userPwd,
		UserName: "",
	}
	//对login结构体进行序列化
	mesInfo, err := json.Marshal(login)
	if err != nil {
		return errors.New("login json Marshal err")
	}
	//完成对mes.Data的赋值
	mes.Data = string(mesInfo)
	//对mes结构体进行序列化
	mesInfo, err = json.Marshal(mes)
	if err != nil {
		return errors.New("mes json Marshal err")
	}
	//在发送数据前 先于服务器确定发送数据的大小
	pkg := uint32(len(mesInfo))
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, pkg)
	//发送校验数据给服务器
	n, err := conn.Write(buf)
	if n != 4 || err != nil {
		return errors.New("conn write err")
	}
	//开始发送mesInfo
	n, err = conn.Write(mesInfo[:pkg])
	if n != int(pkg) || err != nil {
		err = errors.New("conn write err")
		return
	}
	return nil
}
