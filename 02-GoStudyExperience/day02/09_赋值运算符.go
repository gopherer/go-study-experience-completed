package main

import "fmt"

func main09() {
	a := 10

	a += 10
	fmt.Println(a)
	a -= 10
	fmt.Println(a)
	a *= 10
	fmt.Println(a)
	a /= 10
	fmt.Println(a)
	a %= 3
	fmt.Println(a)

	var b = 22
	b %= 2 + 3 //b %= 5 // b = b % 5
	fmt.Println(b)
}
