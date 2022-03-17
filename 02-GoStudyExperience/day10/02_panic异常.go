package main

import "fmt"

func main02() {

	i := 10
	test(i)
	//fmt.Println("hello go")
	//fmt.Println("hello go1")
	//fmt.Println("hello go2")
	////程序遇到panic会终止
	//panic("hello go3")
	//fmt.Println("hello go3")
	//fmt.Println("hello go4")
}
func test(i int) {
	var arr [10]int
	arr[i] = 100
	fmt.Println(arr)
}
