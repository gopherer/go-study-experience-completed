package main

import "fmt"

func main0202() {
	//var a int = 1
	//var b float64 = 3.4
	var (
		a         = 4 //自动推导
		b float64 = 5.5
		c int
	)
	c = 39
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

}
