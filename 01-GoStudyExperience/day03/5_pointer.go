package main

import "fmt"

func swap1(p,q *int) {
	*p,*q = *q ,*p
	fmt.Printf("swap:*p = %d *q = %d\n",*p,*q)
}

func main05() {
	a, b := 10, 20

	swap1(&a, &b)

	fmt.Printf("main : a = %d b = %d\n", a, b)
}
