package main

import (
	"fmt"
	"os"
)

var (
	userID  int
	userPwd string
)

func main() {
	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	var loop bool = true
	//菜单界面
	for loop {
		fmt.Println("-----------------------欢迎登入多人聊天系统-----------------------")
		fmt.Printf("\t\t\t1、用户登入\n")
		fmt.Printf("\t\t\t2、用户注册\n")
		fmt.Printf("\t\t\t3、退出程序\n")
		fmt.Printf("\t\t\t请选择(1-3)\n")
		_, err := fmt.Scanln(&key)
		if err != nil {
			fmt.Println("输入有误")
		}
		switch key {
		//用户登入
		case 1:
			fmt.Println("欢迎进入，用户登入界面")
			loop = false
			//用户注册
		case 2:
			fmt.Println("欢迎进入，用户注册界面")
			loop = false
		case 3:
			fmt.Print("程序已退出")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if key == 1 {
			err = login(userID, userPwd)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("登入成功。。。")
			}
		} else if key == 2 {
			fmt.Println("用户注册界面")
		}

	}
}
