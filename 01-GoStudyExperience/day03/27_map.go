package main

import "fmt"

func main27() {
	m := map[int]string{1: "hello", 2: "world", 3: "go"}
	fmt.Println("m = ", m)

	delete(m, 2) //删除key为1的内容
	fmt.Println("m = ", m)
}
