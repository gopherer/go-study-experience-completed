package main

import "fmt"

func main15() {
	s := "hello"
	//if和{就是条件，条件通常都是关系运算符
	if s == "hello" { //左括号和if在同于行
		fmt.Println("hello world")
	}
	//if支持一个初始化语句，初始化语句和判断语句以分号分隔
	if a := 10; a == 10 {
		fmt.Println("a = ", a)
	}

	var a int
	fmt.Println("请输入a的值：")
	fmt.Scan(&a)

	if a > 10 {
		fmt.Println("a > 10")
	} else if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a < 10")
	}
}
