package main

import "fmt"

func function20(a int) {
	if a == 1 {//函数终止调用的条件，非常重要
		fmt.Println("a = ", a)
		return//终止函数调用
	}
	//函数调用自己本身
	function20(a - 1)
	fmt.Println("a = ", a)
}
func main2108() {
	function20(3)
	fmt.Println("main")
}
