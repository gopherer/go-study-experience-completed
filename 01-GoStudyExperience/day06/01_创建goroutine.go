package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("This is a newTask goroutine")
		time.Sleep(time.Second)
	}
}
func main01() {

	go newTask()
	for {
		fmt.Println("This is a main goroutine")
		time.Sleep(time.Second) //延迟1秒
	}

}
