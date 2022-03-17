package main

import (
	"fmt"
	"sort"
)

func main05() {
	var m = make(map[int]int)
	m[10] = 10
	m[3] = 22
	m[5] = 88
	m[55] = 99
	fmt.Println("map = ", m)

	var slice []int
	for k, _ := range m {
		slice = append(slice, k)
	}
	fmt.Println("slice = ", slice)

	sort.Ints(slice)
	fmt.Println("slice = ", slice)

	for _, v := range slice {
		fmt.Printf("map[%v] = %v\n", v, m[v])
	}
}
