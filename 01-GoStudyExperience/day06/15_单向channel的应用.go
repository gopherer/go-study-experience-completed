package main

import "fmt"

func main15() {
	//创建一个双向通道
	ch := make(chan int)

	//生产者，生产数字，写入channel
	//新建一个协程
	go producer(ch) //channel传参，引用传递，地址传递

	//消费者，从channel读取数据，打印
	consumer(ch)
}
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}
