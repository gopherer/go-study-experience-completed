package main

import "fmt"

func main21() {
	src := []int{1, 2}
	des := []int{6, 6, 6, 6, 6, 6}

	// copy(des, src)
	// fmt.Println("des = ", des)
	// fmt.Println("src = ", src)

	copy(src, des)
	fmt.Println("des = ", des)
	fmt.Println("src = ", src)
	//不会扩容
}
