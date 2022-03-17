package main

import "fmt"

func function31() int {
	//函数被调用是，x,才分配空间，才初始化为0
	var x int //变量没有初始化，默认值为0
	x++
	return x * x ///函数调用完毕，x自动释放
}
func main2114() {
	fmt.Println("x = ", function31())
	fmt.Println("x = ", function31())
	fmt.Println("x = ", function31())
}

//函数的返回值是一个匿名函数，返回一个函数类型
func function32() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main2114_1() {
	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回匿名函数，f来调用闭包函数
	//它不关心这些捕获的变量和常量是否已经超出了作用域
	//所以自由闭包在使用它，这些变量就还存在
	f := function32()
	fmt.Println("f = ", f())
	fmt.Println("f = ", f())
	fmt.Println("f = ", f())
	fmt.Println("f = ", f())
	main2114()
}

//个人理解 闭包中的变量 有点类似与C语言中 函数内的静态变量 但闭包中的变量作用域为全局 (全局静态变量)
