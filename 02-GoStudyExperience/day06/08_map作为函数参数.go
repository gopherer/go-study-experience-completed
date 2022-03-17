package main

import "fmt"

func main08() {
	m := map[int]string{1: "hello", 2: "world", 3: "go"}
	fmt.Printf("%p\n", m)
	demo(m)
	fmt.Println(m)
}
func demo(m map[int]string) {
	fmt.Printf("%p\n", m)
	m[4] = "hahaha"
	fmt.Println(m)
	delete(m, 4)
	fmt.Println(m)
}
