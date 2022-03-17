package main

import "fmt"

type Student6 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test02(p *Student6) {
	p.id = 666
	fmt.Println("test = ", p)
}
func main35() {
	s := Student6{1, "mike", 'm', 18, "fujian"}

	test02(&s) //地址传递(引用传递)，形参可以改变实参
	fmt.Println("main: ", s)
}
