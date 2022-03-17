package main

import (
	"fmt"
	"runtime"
)

func main04() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片，先执行别的协程，执行完毕后，在回来执行
		runtime.Gosched()
		fmt.Println("hello")
	}
}
