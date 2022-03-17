package main

import (
	"fmt"
	"runtime"
)

const (
	num = 100000
)

var CPUNum = runtime.NumCPU()

func main() {
	intChan := make(chan int, num)
	primeChan := make(chan int, num)
	exitChan := make(chan bool, CPUNum)

	go input(intChan)
	for i := 0; i < CPUNum; i++ {
		go prime(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < CPUNum; i++ {
			<-exitChan
		}
		close(exitChan)
		close(primeChan)
	}()
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数有", v)
	}
	fmt.Println("程序已结束.")
}
func input(intChan chan int) {
	for i := 1; i <= num; i++ {
		intChan <- i
	}
	close(intChan)
	fmt.Println("输入已经完成。")
}
func prime(intChan chan int, primeChan chan int, exit chan bool) {
Flag:
	for {
		num, ok := <-intChan
		if !ok {
			fmt.Println("一个协程已结束")
			break
		}
		for i := 2; i < num-1; i++ {
			if num%i == 0 {
				continue Flag
			}
		}
		primeChan <- num
	}
	exit <- true
}
