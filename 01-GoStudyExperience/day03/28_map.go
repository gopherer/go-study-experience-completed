package main

import "fmt"

func test(m map[int]string) {
	delete(m, 2)
}
func main28() {
	m := map[int]string{1: "hello", 2: "world", 3: "go"}
	fmt.Println("m = ", m)

	test(m) //在函数内部删除某个key
	fmt.Println("m = ", m)
	
}
