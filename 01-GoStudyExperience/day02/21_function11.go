package main

import "fmt"

//FuncType1 是函数变量的别名
type FuncType1 func(int, int) int

func function26Add(a, b int) int {
	return a + b
}
func function27Min(a, b int) int {
	return a - b
}
func function28Mul(a, b int) int {
	return a * b
}
func function29Div(a, b int) int {
	return a / b
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同的表现加减乘除
//先有想法，后面在实现功能
func function30Calc(a, b int, fTest FuncType1) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b)
	return

}
func main2111() {
	x, y := 10, 5
	a := function30Calc(x, y, function26Add)
	fmt.Println("a = ", a)
	b := function30Calc(x, y, function27Min)
	fmt.Println("b = ", b)
	c := function30Calc(x, y, function28Mul)
	fmt.Println("c = ", c)
	d := function30Calc(x, y, function29Div)
	fmt.Println("d = ", d)
}
