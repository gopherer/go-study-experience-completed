package main

import "fmt"

func main07() {
	var a [5]int
	fmt.Printf("len(a) = %d\n", len(a))

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}
}
