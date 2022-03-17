package main

import "fmt"

func main03() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	i := 0
	j := len(arr) - 1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println(arr)

	for i := 0; i < len(arr)-i; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	fmt.Println(arr)
}
