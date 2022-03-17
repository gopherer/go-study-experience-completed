package main

import "fmt"

const (
	X = 30
	Y = 3
)

func PreUI() {
	if err := Clear(); err != nil {
	}
	GotPosition(X-5, Y)
	fmt.Print("--在进入系统,请先输入要计算的二进制数--")
	GotPosition(X+5, Y+1)
	fmt.Print("Number1 : ")
	if _, err := fmt.Scan(&user.Number1); err != nil {
	}
	GotPosition(X+5, Y+2)
	fmt.Print("Number2 : ")
	if _, err := fmt.Scan(&user.Number2); err != nil {
	}
}
func UI() int {
	if err := Clear(); err != nil {
	}
	ch := make(chan int, 1)
	for {
		GotPosition(X-5, Y)
		fmt.Print("--欢迎来到二级制计算系统--")
		GotPosition(X, Y+1)
		fmt.Print("->二进制乘法")
		GotPosition(X, Y+2)
		fmt.Print("  退出计算系统")
		HideCursor() //设置控制台光标隐藏
		go func() {
			flag := 1
			for {
				switch Direction() {
				case 13: //回车 Enter
					if Clear() != nil {
					}
					if flag == 1 {
						ch <- 1
					}
					if flag == -1 {
						ch <- -1
					}
				case 72, 87, 119: //↑ W w
					flag = uiMUL()
				case 80, 83, 115: //↓ S s
					flag = uiExit()
				}
			}
		}()
		if <-ch == 1 {
			return 1
		} else if <-ch == -1 {
			return -1
		}
	}
}
func uiMUL() int {
	if Clear() != nil {
	}
	GotPosition(X-5, Y)
	fmt.Print("--欢迎来到二级制计算系统--")
	GotPosition(X, Y+1)
	fmt.Print("->二进制乘法")
	GotPosition(X, Y+2)
	fmt.Print("  退出计算系统")
	return 1
}
func uiExit() int {
	if Clear() != nil {
	}
	GotPosition(X-5, Y)
	fmt.Print("--欢迎来到二级制计算系统--")
	GotPosition(X, Y+1)
	fmt.Print("  二进制乘法")
	GotPosition(X, Y+2)
	fmt.Print("->退出计算系统")
	return -1
}
func (user *User) MulResultUI(result string) {
	GotPosition(X, Y)
	fmt.Print("被除数:", user.Number1)
	GotPosition(X, Y+1)
	fmt.Print("  除数:", user.Number2)
	GotPosition(X, Y+2)
	fmt.Print("  结果:", result)
}
