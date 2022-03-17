package main

import "fmt"

func main10() {
	var a int = 10
	b := 1000000000
	//int会根据操作系统位数的不同在内存中分配的空间也不同
	fmt.Println(a)
	fmt.Println(b)
}

func main1001() {
	var a uint
	a = 10
	fmt.Println(a)
	a = a - 20
	fmt.Println(a)
}
func main1002() {
	var a int = 10
	var b int64 = 28
	//fmt.Println(a+b)//err 类型不匹配 需要类型转换
	fmt.Println(a + int(b))
}
