package main

import "fmt"

//无参有返回值:只有一个返回值
//有返回值的函数需要通过return中断函数，通过return返回
func function10() int {
	return 666
}
//给返回值起一个变量名，go语言推荐写法
//常用写法
func function11() (result int) {
	return 999
}

func function12() (result int) {
	result = 888
	return
}
func main2104() {
	//无参有返回值函数的调用
	var a int
	a = function10()
	fmt.Println("a = ", a)

	b := function11()
	fmt.Println("b = ", b)

	c := function12()
	fmt.Println("c = ", c)
}
