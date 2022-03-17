package main

import (
	"fmt"
	"time"
)

//主协程退出了，其他子协程也会跟着退出
func main02() {
	go func() { //匿名函数
		i := 0
		for {
			i++
			fmt.Println("son goroutine i = ", i)
			time.Sleep(time.Second)
		}
	}() //调用匿名函数

	i := 0
	for {
		i++
		fmt.Println("main goroutine i =", i)
		time.Sleep(time.Second)
		if i == 3 {
			break
		}
	}
}
