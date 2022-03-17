package main

import "fmt"

func main19() {
	arr := []int{}
	fmt.Printf("len = %d cap = %d\n", len(arr), cap(arr))
	fmt.Println("arr = ", arr)
	//append函数在原切片末尾追加元素
	arr = append(arr, 888)
	arr = append(arr, 999)
	fmt.Printf("len = %d cap = %d\n", len(arr), cap(arr))
	fmt.Println("arr = ", arr)

	s1 := []int{1, 2, 3}
	fmt.Println("s1 = ", s1)
	s1 = append(s1, 555)
	s1 = append(s1, 999)
	fmt.Printf("len = %d cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

}
