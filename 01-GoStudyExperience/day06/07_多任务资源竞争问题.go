package main

import (
	"fmt"
	"time"
)

//Printer 打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	//fmt.Println("")
}
func person1() {
	Printer("hello")
}
func person2() {
	Printer("world")
}
func main07() {
	//新建2个协程，代表2个人，2个人同时使用打印机
	go person1()
	go person2()

	//主协程，死循环，保证子协程的运行
	for {

	}
}
