package main

import "fmt"

type Student5 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test01(s Student5) {
	s.id = 666
	fmt.Println("test = ", s)
}
func main34() {
	s := Student5{1, "mike", 'm', 18, "fujian"}

	test01(s)//值传递，形参无法改变实参
	fmt.Println("main: ", s)
}
