package main

import "fmt"

func main08() {
	//1、全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)
	//部分初始化，为初始话部分自动赋值为0
	c := [5]int{1, 3, 4}
	fmt.Println("c = ", c)
	//指定某个元素初始化
	d := [5]int{2: 10, 4: 70}
	fmt.Println("d = ", d)

}
