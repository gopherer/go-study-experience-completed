package main

import "fmt"

type Student1 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main30() {
	//顺序初始化，每个成员必须初始化 别忘了&
	var p *Student1 = &Student1{1, "mike", 'm', 18, "fujian"}
	fmt.Println("*p = ", *p)
	//指定成员初始化，没有初始化的成员，自动赋值为0
	p2 := &Student1{id: 2, name: "hahaha"}
	fmt.Println("p2 = ", p2)
}
