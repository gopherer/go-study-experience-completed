package main

import "fmt"

func main11() {
	var a bool
	fmt.Println(!a)

	b := true

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(a && b || a)

}
