package main

import "fmt"

func main13() {
	ch := make(chan int)

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//不需要再写数据时，关闭channel
		close(ch)
		// ch <- 666//panic 关闭channel后无法在发送数据
	}()

	for num := range ch { //通过迭代实现ch管道的遍历
		fmt.Println("num = ", num)
	}
}
