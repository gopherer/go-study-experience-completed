package main

import "fmt"

func main() {
	arr := [5]int{20, 10, 59, 80, 30}
	fmt.Println("old arr = ", arr)
	selectSort(&arr)
	fmt.Println("new arr = ", arr)
}
func selectSort(arr *[5]int) {
	/*
		tem := arr[0]
		index := 0
		for i := 0 + 1; i < len(arr); i++ {
			if tem < arr[i] {
				tem = arr[i]
				index = i
			}
		}
		if index != 0 {
			arr[0], arr[index] = arr[index], arr[0]
		}
	*/
	for j := 0; j < len(arr)-1; j++ {
		tem := arr[j]
		index := j
		for i := j + 1; i < len(arr); i++ {
			if tem < arr[i] {
				tem = arr[i]
				index = i
			}
		}
		if index != j {
			arr[j], arr[index] = arr[index], arr[j]
		}
	}
}
