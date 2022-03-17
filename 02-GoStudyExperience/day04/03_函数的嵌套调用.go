package main

import "fmt"

func test1(a, b int) {
	fmt.Println(a + b)
}
func test(a int, b int) {
	test1(a, b)
}
func main03() {
	a := 10
	b := 20
	test(a, b)
}

func main0301() {
	test2(1, 2, 3, 4, 5)
}
func test2(arr ...int) {
	test3(arr)
}
func test3(arr []int) {
	fmt.Println(arr)
}
