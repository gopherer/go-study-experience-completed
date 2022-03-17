package main

import "fmt"

func main02() {
	slice := []int{1, 9, 8, 4, 5, 6, 3, 2, 7, 10}
	fmt.Printf("%p\n", slice)
	//切片作为函数参数 地址传递 引用传递
	test(slice)
	fmt.Println(slice)
	slice = test1(slice)
	fmt.Println(slice)
}

func test(s []int) {
	fmt.Printf("%p\n", s)
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func test1(s []int) []int {
	fmt.Printf("%p\n", s)
	//使用 append 会改变切片的地址
	s = append(s, 100, 200, 300)
	fmt.Printf("%p\n", s)
	return s
}
