package main

import "fmt"

func main() {
	arr := [6]int{28, 90, 8, 38, 28, 83}
	fmt.Println("old arr = ", arr)
	insertSort(&arr)
	fmt.Println("new arr = ", arr)
}

func insertSort(arr *[6]int) {
	/*
		insertVal := arr[1]
		insertIndex := 1 - 1

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != 1 {
			arr[insertIndex+1] = insertVal
		}
	*/
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}
