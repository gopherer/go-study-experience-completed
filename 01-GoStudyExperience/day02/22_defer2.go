package main

import "fmt"

func main2202_1() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("a = %d b = %d\n", a, b)
	}(a, b) //()代表调用此匿名函数，已经先传递了参数，只是没有调用

	a = 111
	b = 222
	fmt.Printf("a = %d b = %d\n", a, b)
}

func main2202() {
	a := 10
	b := 20

	defer func() {
		fmt.Printf("a = %d b = %d\n", a, b)
	}() //()代表调用此匿名函数

	a = 111
	b = 222
	fmt.Printf("a = %d b = %d\n", a, b)
}
