package main

import "fmt"

func main04() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan type  %T  content %v addr %v\n", intChan, intChan, &intChan)
	intChan <- 10
	var num = 30
	intChan <- num
	fmt.Printf("intChan len %v cap %v\n", len(intChan), cap(intChan))
	num1 := <-intChan
	fmt.Println(num1)
}
