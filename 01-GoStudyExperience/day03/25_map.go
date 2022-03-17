package main

import "fmt"

func main25() {
	m1 := map[int]string{1: "hello", 2: "world"}
	fmt.Println("m1 = ", m1)
	//赋值，如果已经存在key值，修改内容
	m1[1] = "go"
	m1[3] = "c" //追加，map底层自动扩容，与append类似
	fmt.Println("m1 = ", m1)
}
