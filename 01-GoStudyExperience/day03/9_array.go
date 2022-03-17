package main

import "fmt"

func main09() {
	var a [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d", i, j, a[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("a = ", a)
	//有3个元素，每个元素又是一个数组[4]int
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b = ", b)
	//部分初始化，未初始化的值为0
	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9}}
	fmt.Println("c = ", c)

	d := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println("d = ", d)
	//选择性初始化，未初始化值为0
	e := [3][4]int{2: {1, 3, 4, 5}}
	fmt.Println("e = ", e)
}
