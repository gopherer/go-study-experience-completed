package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap:a = %d b = %d\n", a, b)
}

func main04() {
	a, b := 10, 20

	swap(a, b)

	fmt.Printf("main : a = %d b = %d\n", a, b)
}
