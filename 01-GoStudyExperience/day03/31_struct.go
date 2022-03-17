package main

import "fmt"

type Student2 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main31() {
	//定义一个结构体变量
	var s Student2

	//操作结构体成员，需要使用点(.)运算符
	s.id = 1
	s.name = "mike"
	s.age = 19
	fmt.Println("s = ", s)
}
