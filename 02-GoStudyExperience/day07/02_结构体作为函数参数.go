package main

import "fmt"

func test(stu student) {
	stu.age = 100
	fmt.Println(stu)
}
func main02() {
	stu := student{"李四", 18, 30}
	//变量--值传递
	test(stu)
	fmt.Println(stu)
}

func main0201() {
	m := map[int]student{
		100: {name: "张三", age: 19, score: 100},
		200: {"小明", 20, 100},
	}
	//地址传递 引用传递
	demo(m)
	fmt.Println(m)
}

func demo(m map[int]student) {
	//m[100].age = 100 err
	temp := m[100]
	temp.age = 100
	m[100] = temp
	fmt.Println(m)
}
