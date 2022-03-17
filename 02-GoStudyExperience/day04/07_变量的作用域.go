package main

import "fmt"

var a int = 100

func main07() {
	//var i int = 10
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println(i)
	//
	//var a int = 10
	//for a = 0; a < 5; a++ {
	//	fmt.Println(a)
	//}
	//fmt.Println(a)
	//a = 999
	a := 777
	fmt.Println(a)
	demo3()
}

func demo3() {
	fmt.Println(a)
}

func main0701() {
	fmt.Println(demo3)
	fmt.Println(&a)
	a := 99
	fmt.Println(&a)
}
