package main

import "fmt"

func main() {
	arr := [10]int{5, 4, 3, 9, 7, 2, 3, 8, 2, 6}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
