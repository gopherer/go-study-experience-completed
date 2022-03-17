package main

import "fmt"

func main06() {
	a := 10
	b := 20
	c := 30

	var arr [3]*int = [3]*int{&a, &b, &c}

	fmt.Println(arr)
	fmt.Println(&a, &b, &c)

	*arr[0] = 100
	fmt.Println(a)
}

func main0601() {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	c := [3]int{7, 8, 9}

	var arr [3]*[3]int = [3]*[3]int{&a, &b, &c}
	fmt.Println(arr)

	fmt.Printf("arr type is %T\n", arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", (*arr[i])[j])
		}
	}
	fmt.Println("")
	for i := 0; i < len(arr); i++ {
		fmt.Println(*arr[i])
	}
}
