package main

import "fmt"

//形成闭包  个人认为 变量内存逃逸
func addUpper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(i int) int {
		n += i
		str += string(i) //36 对应的assic码为$
		fmt.Println(str)
		return n
	}
}
func main03() {
	f := addUpper()
	fmt.Println(f(36))
	fmt.Println(f(36))
}
