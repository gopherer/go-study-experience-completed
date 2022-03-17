package main

import (
	"fmt"
	"time"
)

func main19() {
	timer := time.NewTimer(3 * time.Second)
	timer.Reset(1 * time.Second)

	<-timer.C
	fmt.Println("时间到")
}
