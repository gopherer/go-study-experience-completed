package main

import "fmt"

func sum(arr ...int) {
	fmt.Println(arr)
	fmt.Printf("arr type = %T\n", arr)

	fmt.Println(arr[0])
	fmt.Println(arr[1])

	count := len(arr)
	fmt.Println(count)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, value := range arr {
		fmt.Println(value)
	}

}
func main02() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
