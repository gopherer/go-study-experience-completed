package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaa")
}
func testb() {
	//fmt.Println("bbbbbbbbbbb")
	//显式调用panic函数，导致程序中断
	panic("this is a panic test")
}
func testc() {
	fmt.Println("ccccccccccc")
}
func main03() {
	testa()
	testb()
	testc()
}
