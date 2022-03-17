package main

import "fmt"

func main09() {
	//while
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello", i)
		i++
	}
	fmt.Println("i = ", i)
	//do...while
	var j int = 1
	for {
		fmt.Println("hello ok", j)
		j++
		if j > 10 {
			break
		}
	}
	fmt.Println("j = ", j)
}
