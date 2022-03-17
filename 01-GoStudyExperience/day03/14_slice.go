package main

import "fmt"

func main14() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:3:5] //a[起始下标：终止下标(不包括):最大值]
	fmt.Println("s = ", s)
	fmt.Println("lem(s) = ", len(s)) //长度 3-0
	fmt.Println("cap(s) = ", cap(s)) //容量 5-0

	s = a[1:4:5]
	fmt.Println("s = ", s)           //从下标1开始，取4-1=3个
	fmt.Println("lem(s) = ", len(s)) //长度 4-1
	fmt.Println("cap(s) = ", cap(s)) //容量 5-1

}
