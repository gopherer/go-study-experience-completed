package main

import "fmt"

func main24() {
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len,没有cap
	fmt.Println("len = ", len(m1))
	//通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))
	//可以通过make创建，可以指定长度，只是指定容量，但是里面一个数据都没有
	m3 := make(map[int]string, 2)
	m3[1] = "hello"
	m3[2] = "world"
	m3[666] = "go"
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))
	//初始化
	//键值唯一
	m4 := map[int]string{1: "hhh", 2: "jjj"}
	fmt.Println("m4 = ", m4)
	fmt.Println("len = ", len(m4))
}
