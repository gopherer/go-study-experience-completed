package main

import "fmt"

func function19(c int) {
	fmt.Println("c = ", c)
}
func function18(b int) {
	function19(b - 1)
	fmt.Println("b = ", b)
}
func function17(a int) {
	function18(a - 1)
	fmt.Println("a = ", a)
}
func main2107() {
	function17(4)
	fmt.Println("main")
}

//函数调用流程，先调用后返回，先进后出
//函数递归，函数调用自己本身，利用此特点
