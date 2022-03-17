package main

import "fmt"

func testaa() {
	fmt.Println("aaaaaaaaaaaaaa")
}
func testbb(x int) {
	var a [10]int
	a[x] = 111
}
func testcc() {
	fmt.Println("ccccccccccc")
}
func main04() {
	testaa()
	testbb(20)
	testcc()
}