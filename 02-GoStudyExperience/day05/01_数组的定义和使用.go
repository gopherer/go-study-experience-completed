package main

import "fmt"

func main01() {
	var arr [10]int

	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[5])
}
func main0101() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr[0] = 100
	fmt.Println(arr)

	count := len(arr)
	fmt.Println(count)

	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
		fmt.Printf("arr[%d] = %d ", i, arr[i])
	}

	for i, value := range arr {
		fmt.Println(i, value)
	}
}

func main0102() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	var arr1 [10]int
	arr1 = arr
	fmt.Println(arr1)

	fmt.Printf("arr type is %T\n", arr)

	fmt.Printf("arr地址 %p\n", &arr)
	fmt.Printf("arr[0]地址 %p\n", &arr[0])
	fmt.Printf("arr[1]地址 %p\n", &arr[1])

}
