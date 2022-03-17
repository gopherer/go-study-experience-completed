package main

import "fmt"

func main01() {
	a := 10
	b := 20
	add(a, b)
}
func add(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}
