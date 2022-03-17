package main

import (
	"fmt"
	"time"
)

//主协程退出了，其他子协程也会跟着退出
//主协程先退出导致子协程未调用完全就退出
func main03() {
	go func() { //匿名函数
		i := 0
		for {
			i++
			fmt.Println("son goroutine i = ", i)
			time.Sleep(time.Second)
		}
	}() //调用匿名函数
}
