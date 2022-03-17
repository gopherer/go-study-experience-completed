package main

import (
	"fmt"
	"unsafe"
)

func array(arr *[3]int) {
	(*arr)[0] = 10
	(*arr)[1] = 20
	fmt.Println("*arr = ", *arr, " arr = ", arr)
	fmt.Printf("arr size is %d\n", unsafe.Sizeof(arr))
	fmt.Printf("*arr size is %d\n", unsafe.Sizeof(*arr))
	fmt.Printf("arr type is %T\n", arr)
}
func main08() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)
	array(&arr)
	fmt.Println(arr)

}
