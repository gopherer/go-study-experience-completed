package main

import "fmt"

func test4(a, b int) int {
	sum := a + b
	return sum
}
func main05() {
	a := 10
	b := 20
	sum := test4(a, b)
	fmt.Println(sum)
}

func test5(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
func main0501() {
	a := 10
	b := 20
	sum, sub := test5(a, b)
	fmt.Println(sum, sub)
}
