package main

import "fmt"

func main07() {
	slice := []int{1, 2, 3, 4, 5}
	p := &slice

	fmt.Printf("%T\n", p)
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", slice)

	//p[1] = 222 err
	(*p)[1] = 222
	fmt.Println(slice)

	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
}

func main0701() {
	slice := []int{1, 2, 3}

	demo1(slice)
	fmt.Println(slice)

	demo2(&slice)
	fmt.Println(slice)
}

func demo1(slice []int) {
	slice = append(slice, 12, 3, 4)
}

func demo2(slice *[]int) {
	*slice = append(*slice, 12, 3, 4)
}
