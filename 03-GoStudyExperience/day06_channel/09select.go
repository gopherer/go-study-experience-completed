package main

import (
	"fmt"
	"strconv"
)

func main09() {
	c := make(chan int, 10)
	d := make(chan string, 5)

	for i := 0; i < cap(c); i++ {
		c <- i
	}
	for i := 0; i < cap(d); i++ {
		d <- strconv.Itoa(i) + "hello"
	}
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		case j := <-d:
			fmt.Println(j)
		default:
			return
		}

	}
}
