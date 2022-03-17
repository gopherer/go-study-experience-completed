package main

import "fmt"

func main05() {
	a,b:= 10,20
	fmt.Println(a,b)

	a,b = b,a
	fmt.Println(a,b)
}
