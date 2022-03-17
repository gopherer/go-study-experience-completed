package main

import "fmt"

func main11() {
	slice := []int{1, 2, 3, 4, 5}

	s := make([]int, 3)

	count := copy(s, slice)
	fmt.Println(count)
	fmt.Println(cap(s))
	fmt.Println(s)
	fmt.Println(slice)

	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", s)

}
