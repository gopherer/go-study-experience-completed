package main

import (
	"fmt"
	"time"
)

func main18() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("定时器时间到，子协程打印")
	}()
	timer.Stop() //停止定时器
	for {

	}
}
