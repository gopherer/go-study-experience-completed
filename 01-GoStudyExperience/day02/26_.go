package main

import "fmt"

//var a byte //全局变量

func main26() {
	var a int //局部变量

	//1、不同作用域，允许定义同名变量
	//2、使用变量的原则，就近原则
	fmt.Printf("%T\n", a) //int
	{
		var a float64
		fmt.Printf("%T\n", a) //float64
	}
	test02()
}
func test02() {
	fmt.Printf("%T\n", a) //uint8
}
