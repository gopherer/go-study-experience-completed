package main

import (
	"fmt"
)

func Write(intChan chan<- int) {
	for i := 0; i < 20; i++ {
		intChan <- i
		fmt.Println("write i = ", i)
	}
	close(intChan)
}
func Read(intChan <-chan int, exitChan chan<- bool) {
	for {
		i, ok := <-intChan
		if !ok {
			fmt.Println(i, ok)
			break
		}
		fmt.Println("read i = ", i)
	}
	exitChan <- true
	close(exitChan)
}
func main07() {
	intChan := make(chan int, 20)
	exitChan := make(chan bool, 1)
	go Write(intChan)
	go Read(intChan, exitChan)
	//<-exitChan //阻塞等待
	//或使用
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
