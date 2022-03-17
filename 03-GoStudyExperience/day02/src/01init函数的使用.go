package main

import (
	"demo"
	"fmt"
)

//全局变量会最先被执行
var age int = test()

func test() int {
	fmt.Println("hello go")
	return 100
}

//init 函数会在main函数之前被go框架调用 通常用于初始化
func init() {
	fmt.Println("test ...")
}
func main() {
	fmt.Println("main...")
	fmt.Println("age = ", age)
	fmt.Println("demo.Name = ", demo.Name)
	fmt.Println("demo.Age = ", demo.Age)

}
