package main

import (
	"fmt"
)

func main09() {
	var slice []int
	fmt.Printf("%p\n", slice) //空切片 指向内存地址位0的空间
	slice = append(slice, 1, 2, 3)
	fmt.Printf("%p\n", slice)
	//append 在切片末尾追加内容 切片地址可能发生改变
	slice = append(slice, 4, 5, 6)
	fmt.Printf("%p\n", slice)
}
func main0901() {
	slice := make([]int, 10, 20)
	//扩容过程一般位上一次容量的两倍 容量超过1024byte 一般扩容位上一次的1/4
	//扩容后 容量一般位偶数
	//扩容是可能会改变切片的地址
	fmt.Printf("%p\n", slice)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("%p\n", slice)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("%p\n", slice)
	fmt.Println(slice, len(slice), cap(slice))
}
