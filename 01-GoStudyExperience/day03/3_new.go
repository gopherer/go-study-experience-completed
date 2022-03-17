package main

import "fmt"

func main03() {
	var p *int
	p = new(int)
	*p = 666
	fmt.Println("*p = ", *p)
	//new()相当于c语言中动态分配内存空间 但无需手动释放
	q := new(int)
	*q = 999
	fmt.Println("*q = ", *q)
}
