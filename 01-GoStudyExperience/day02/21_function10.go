package main

import "fmt"

func function24(a int, b int) int {
	return a + b
}
func function25(a int, b int) int {
	return a - b
}

//FuncType 是函数变量的别名 注意没有函数名字，没有{}
//函数也是一种数据类型，通过type给函数类型起一个别名
//FuncType它是一个数字类型
type FuncType func(int, int) int

func main2110() {
	var result int
	result = function24(2, 9) //传统调用方法
	fmt.Println("result = ", result)

	//声明一个函数类型的变量，变量名为fTest
	var fTest FuncType
	fTest = function24 //是变量就可以赋值
	result = fTest(2, 8)//等价于function24(2,8)
	fmt.Println("result = ", result)

	var test func(int, int) int
	test = function25
	result = test(5, 2)//等价于function25(5,2)
	fmt.Println("result = ", result)

}
