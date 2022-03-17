package main

import "fmt"

func main18() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//新切片
	s := a[2:5]
	s[1] = 555
	fmt.Println("s = ", s)
	fmt.Println("a = ", a)

	s1 := a[1:6]
	s1[3] = 888
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)
}
