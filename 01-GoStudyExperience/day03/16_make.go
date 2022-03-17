package main

import "fmt"

func main16() {
	s := []int{1, 2, 3, 4}
	fmt.Println("s = ", s)
	//借助make函数，格式 make(切片类型，长度，容量)
	s1 := make([]int, 5, 10)
	fmt.Printf("len = %d cap = %d\n", len(s1), cap(s1))
	//没有指定容量，容量和长度一样
	s2 := make([]int, 5)
	fmt.Printf("len = %d cap = %d\n", len(s2), cap(s2))
}
