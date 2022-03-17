package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main29() {
	//顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "mike", 'm', 18, "fujian"}
	fmt.Println("s1 = ", s1)
	//指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := Student{id: 2, name: "hahaha"}
	fmt.Println("s2 = ", s2)
}
