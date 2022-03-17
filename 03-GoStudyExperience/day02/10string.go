package main

import "fmt"

func demo(str string) {
	fmt.Println("str = ", str)
	str = "go"
	fmt.Println("str- = ", str)
}
func main10() {
	var str string = "hello"
	demo(str) //值传递
	fmt.Println("str = ", str)
	//字符串 底层是一个byte的数组
	slice := str[:]
	fmt.Println("slice = ", slice)
}
