package main

import "fmt"

func main15() {
	//切片与数组的区别
	//数组[]里面的长度是一个固定的常量，数组不能修改长度，len和cap的值一样
	a := [5]int{}
	fmt.Printf("len = %d cap = %d\n", len(a), cap(a))

	//切片，[]里面为空，或者... 切片的长度可以不固定 可以修改
	s := []int{}
	fmt.Printf("1: len = %d cap = %d\n", len(s), cap(s))

	s = append(s, 11) //通过append给切片末尾追加一个成员
	fmt.Printf("append: len = %d cap = %d\n", len(s), cap(s))
	fmt.Println("s = ", s)
}
