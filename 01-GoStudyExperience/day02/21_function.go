package main

import "fmt"

//无参无返回值函数的定义
func function() {

	a := 8
	fmt.Println("a = ", a)
}

func main21() {

	function()
	function01()

}

//无需声明即可调用
func function01() {
	a := 9
	fmt.Println("a = ", a)
}
