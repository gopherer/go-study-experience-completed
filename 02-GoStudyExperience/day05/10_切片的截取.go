package main

import "fmt"

func main10() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}

	s := slice[1:3]

	fmt.Println(slice)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	s1 := slice[2:]
	fmt.Println(s1)

	s2 := slice[:2]
	fmt.Println(s2)

	s3 := slice[1:3:5]
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))
}
func main1001() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}

	s := slice[:] //等价与 s = slice
	fmt.Println(s, slice)

	s1 := slice[1:3]

	s1[1] = 100
	fmt.Println(s1, slice)

	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", s1)
}
