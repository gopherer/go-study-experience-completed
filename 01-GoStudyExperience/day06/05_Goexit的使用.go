package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccc")
	//return //结束当前函数
	runtime.Goexit() //结束所在协程
	fmt.Println("dddddddddddddd")
}
func main05() {
	//创建一个新的协程
	go func() {
		fmt.Println("aaaaaaaaaaaaaa")

		//调用别的函数
		test()

		fmt.Println("bbbbbbbbbbbbbbbb")
	}()

	//特意写个死循环，目的为了不让主协程结束
	for {
	}
}
