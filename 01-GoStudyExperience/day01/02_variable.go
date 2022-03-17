package main

import "fmt" //导入包 必须使用

func main() {
	//1、声明格式 var 变量名 类型， 变量声明必须使用
	//2、只有声明没有初始化的变量，默认值为0
	//3、同一个{ }里，声明的变量名是唯一的
	var a int
	fmt.Println("a = ", a)
	//4、可以同时声明多个变量
	//var b, c int
	a = 10 //变量的赋值
	fmt.Println("a = ", a)
	//5、变量的初始化，变量声明时，同时赋值
	var b int = 20 //初始化，声明变量时，同时赋值（一步到位）
	fmt.Println("b = ", b)
	//6、自动推导类型，必须初始化，通过初始化的值来确定类型（常用）
	c := 30
	fmt.Printf("c type is %T\n", c) //%T 打印变量所属的类型
	fmt.Println("c = ", c)
}
