package main

import "fmt"

//有参无返回值函数的定义，普通参数列表
//定义函数时，在函数名后面()定义的参数叫形参
//参数传递，只能由实参传递给形参，不能反过来单向传递
func function02(a int) {
	//a = 644
	fmt.Println("a = ", a)
}

func function03(a int, b int) {
	fmt.Printf("a = %d b = %d\n", a, b)
}

func function04(a int, c string) {
	fmt.Printf("a = %d c = %s\n", a, c)
}

func main2101() {
	//有参无返回值函数调用：函数名（所需参数）
	//调用函数传递的参数叫实参
	function02(666)
	function03(555, 999)
	function04(6666, "ffff")
}
