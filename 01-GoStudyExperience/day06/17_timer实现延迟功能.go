package main

import (
	"fmt"
	"time"
)

func main17() {
	//延迟2s后打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}

func main1701() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

func main1702() {
	<-time.After(2 * time.Second)
	fmt.Println("时间到") //定时2s,阻塞2s,2s后产生一个时间，往channel写内容
}
