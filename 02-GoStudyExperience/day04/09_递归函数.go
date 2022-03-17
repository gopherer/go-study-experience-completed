package main

import "fmt"

func demo4(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	demo4(n - 1)
}

func main09() {
	demo4(10)
}

//s:=1 err
var s int = 1

func demo5(n int) {
	if n == 0 {
		return
	}
	s *= n
	demo5(n - 1)
}
func main0901() {
	demo5(10)
	fmt.Println(s)
}
