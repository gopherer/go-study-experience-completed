package main

import "fmt"

func main04() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
