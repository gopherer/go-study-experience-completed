package main

import "fmt"

func main03() {
	fmt.Println("喜羊羊")
	defer fmt.Println("懒羊羊")
	defer fmt.Println("美羊羊")
	fmt.Println("沸羊羊")
}
func main0301() {
	a := 10
	b := 20

	f := func(a int, b int) {
		fmt.Println(a + b)
	}
	//defer 调用函数如果有传递参数会在内存中存储 外部变量改变不会影响函数的值
	defer f(a, b)

	a = 123
	b = 321
}
func main0302() {
	a := 10
	b := 20

	f := func() {
		fmt.Println(a + b)
	}
	//defer 调用函数如果没有传递参数会 外部变量改变会影响函数的值
	defer f()

	a = 123
	b = 321
}
