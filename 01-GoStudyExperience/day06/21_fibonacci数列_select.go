package main

import "fmt"

func main21() {
	ch := make(chan int)    //数字通道
	quit := make(chan bool) //判断程序是否结束

	//消费者，从channel读取数据
	//新建协程
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		//可以停止
		quit <- true
	}()

	//生产者，产生数字，写入channel
	fibonacci(ch, quit)
}

//ch只写 quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//select 监听channel数据流
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}
