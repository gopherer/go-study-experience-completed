package main

import "fmt"

func main08() {
	var slice []int = []int{1, 2, 3, 4, 5}
	//slice[5] = 6 err index out of range
	slice[1] = 123
	fmt.Println(slice)

	slice = append(slice, 1)
	slice = append(slice, 2, 3, 4, 5)
	fmt.Println(slice)
}
func main0801() {
	slice := []int{1, 2, 3}

	fmt.Println(len(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func main0802() {
	slice := make([]int, 10)
	fmt.Println(len(slice))
	fmt.Println(slice)

	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(slice)
}
