package main

import "fmt"

func function16(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return//有返回值的函数，必须通过return返回
}
func main2106() {
	max, min := function16(29, 77)
	fmt.Printf("max = %d min = %d\n", max, min)
	//匿名变量丢弃某个返回值
	a , _ := function16(29, 77)
	fmt.Printf("a = %d ", a)
}
