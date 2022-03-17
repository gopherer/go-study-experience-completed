package main

import "fmt"

func main02() {
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println(max)
}
