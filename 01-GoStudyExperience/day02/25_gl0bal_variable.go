package main 

import "fmt"

var a int
//定义在函数之外的变量就是全局变量
//全局变量在任何地方都可以使用
func test2(){
	fmt.Println("a = ",a)
}
func main25(){
	a = 9
	fmt.Println("a = ",a)
	test2()
}