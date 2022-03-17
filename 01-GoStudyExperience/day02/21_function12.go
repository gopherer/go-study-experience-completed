package main

import "fmt"

func main2112() {
	a := 10
	str := "mike"

	//匿名函数，没有函数名，
	//函数定义，还未调用
	f1 := func() { //:= 自动推导类型
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	f1()

	//给函数类型起一个别名
	type FuncType2 func() //函数没有参数，没有返回值
	//声明变量
	var f2 FuncType2
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Printf("a = %d str = %s\n", a, str)
	}() //后面的()就代表此匿名函数

	//带参数的匿名函数
	func(i, j int) {
		fmt.Println("i = ", i, "j = ", j)
	}(2, 9)

	//定义匿名函数，并同时调用
	//匿名函数，有参有返回值
	x, y := func(x, y int) (max, min int) {
		if x > y {
			max = x
			min = y
		} else {
			max = y
			min = x
		}
		return
	}(28, 98)
	fmt.Printf("x = %d ,y = %d", x, y)
}
