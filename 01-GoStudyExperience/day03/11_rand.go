package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子，只需一次
	//如果种子的参数一样，那么每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano()) //以当前系统时间为种子参数

	for i := 0; i < 5; i++ {
		//fmt.Println("rand = ", rand.Int())随机出很大的数
		fmt.Println("rand = ", rand.Intn(100)) //随机出的数限制在100以内
	}
}
