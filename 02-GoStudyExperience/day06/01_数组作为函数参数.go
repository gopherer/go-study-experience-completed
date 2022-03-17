package main

import "fmt"

func BubbleSort(arr [10]int) [10]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main01() {
	var arr [10]int = [10]int{2, 6, 9, 8, 4, 1, 5, 4, 0, 7}
	//数值作为函数参数 --- 值传递
	arr = BubbleSort(arr)
	fmt.Println(arr)
}
