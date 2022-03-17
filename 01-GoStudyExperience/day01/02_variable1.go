package main

import "fmt" //导入包 必须使用

func main0201() {
	//赋值，赋值前，必须先声明变量
	var a int
	a = 10
	a = 20
	a = 30 //赋值可以使用多次
	fmt.Println("a = ", a)

	// := ,自动推导类型，先声明变量b，在给b赋值为20
	//自动推导，同一个变量名只能使用一次，用于初始化那次
	b := 20
	fmt.Println("b = ", b)

	//b := 30 //迁移已有变量b,不能再新建一个变脸b

	b = 30 //只有赋值，可以
	fmt.Println("b = ", b)
}
