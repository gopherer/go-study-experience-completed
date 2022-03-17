package main

import "fmt"

func main06() {
	var i interface{}

	fmt.Printf("i type is %T\n", i)
	i = 10
	fmt.Printf("i type is %T\n", i)
	fmt.Println(i)
	i = .15 //i = 0.15
	fmt.Printf("i type is %T\n", i)
	fmt.Println(i)
	i = "hello go"
	fmt.Printf("i type is %T\n", i)
	fmt.Println(i)
	//接口类型不能直接转换为其他数据类型 转换需要类型断言
}
func test() {
	fmt.Println("hello world")
}
func main0601() {
	var i []interface{}

	i = append(i, 110, 2.33, "hello go", test)
	for j := 0; j < len(i); j++ {
		fmt.Println(i[j])
	}
}
