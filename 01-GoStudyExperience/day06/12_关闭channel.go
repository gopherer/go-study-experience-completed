package main

import "fmt"

func main12() {
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

	for {
		//如果ok为true,说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else { //ok ==false 管道关闭
			break
		}
	}

}