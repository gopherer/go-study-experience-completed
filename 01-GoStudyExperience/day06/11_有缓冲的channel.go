package main

import (
	"fmt"
	"time"
)

func main11() {
	ch := make(chan int, 3)
	//len(ch)代表缓冲区c存在数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d cap(ch) = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("子协程: i = ", i)
			fmt.Printf("len(ch) = %d cap(ch) = %d\n", len(ch), cap(ch))
		}
	}()

	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("num = ", num)
	}

}
