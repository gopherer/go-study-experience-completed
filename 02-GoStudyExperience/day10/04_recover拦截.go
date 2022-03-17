package main

import "fmt"

func demo(i int) {

	//recover需要放在函数中 配合defers使用
	defer func() {
		//recover 错误拦截 拦截panic异常
		err := recover()
		fmt.Println(err)
	}()
	var p *int
	*p = 222

	var a [10]int
	a[i] = 100
	fmt.Println(a)
}
func main04() {
	demo(10)
}
