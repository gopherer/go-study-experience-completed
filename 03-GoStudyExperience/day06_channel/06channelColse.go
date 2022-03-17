package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 3
	intChan <- 2
	close(intChan)
	//intChan <- 1
	//channel关闭后不允许写入数据 但可以读取数据
	//channel的遍历
	intChan2 := make(chan int, 20)
	for i := 0; i < 20; i++ {
		intChan2 <- i
	}
	//for i := 0; i < len(intChan2); i++ { //用len(intChan2)只能遍历一般intChan2 可使用cap 前提channel 必须满
	//fmt.Println(<-intChan2)
	//}
	//在用range 需要close()
	close(intChan2)
	for v := range intChan2 {
		fmt.Println(v)
	}
}

/*
当最后一个值从已关闭的clannle中被接受后返回 false，可以用来判断channle执行完成。
package main

import (
	"fmt"
)

func main() {

	var chan1 = make(chan int, 2)

	chan1 <- 10
	chan1 <- 20
	close(chan1)
	<-chan1
	<-chan1

	v, ok := <-chan1
	fmt.Println(v, ok)
}
*/
