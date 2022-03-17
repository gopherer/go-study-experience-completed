package main

import "fmt"

func main12() {
	slice := []int{1, 4, 5, 6, 2, 3, 9, 8, 7, 10}

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)
}
