package main

import (
	"fmt"
	"time"
)

func main16() {
	//创建一个计时器，设置时间为2s，2s后，往time通道写内容(当前时间)
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间: ", time.Now())

	//2s后,往time.C写数据，有数据后就可以读取
	t := <-timer.C //channel没有数据前会阻塞
	fmt.Println("t = ", t)
}

//验证time.NewTimer()，时间到了，只会响应一次
func main1601() {
	{
		time := time.NewTimer(1 * time.Second)
		for {
			<-time.C
			fmt.Println("时间到")
		}
	}
}
