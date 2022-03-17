package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个channel
var ch = make(chan int) //管道默认没有内容

//printer 定义一个打印机 参数为字符串，打印每个字符 属于公共资源
func printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	//fmt.Println("")
}

func hunman1() {
	printer("hello")
	ch <- 666 //给管道写数据，发送数据
}
func hunman2() {
	<-ch //读取，接收管道数据，如果管道没有数据就是堵塞，--不执行后面的语句
	printer("world")
}
func main08() {
	//新建2个协程，代表2个人，2个人同时使用打印机
	go hunman1()
	go hunman2()

	//主协程，死循环，保证子协程的运行
	for {

	}
}
