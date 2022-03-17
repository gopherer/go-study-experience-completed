package main

import (
	"fmt"
	"time"
)

func main19() {
	i := 0
	for { //for循环后面不写任何东西，这个循环条件永远为真，死循环
		i++
		time.Sleep(time.Second) //睡眠一秒
		if i == 4 {
			//break
			continue //只能用在for
		}
		if i == 9 {
			break //可用于for switch select
		}
		fmt.Println("i = ", i)
	}
}
