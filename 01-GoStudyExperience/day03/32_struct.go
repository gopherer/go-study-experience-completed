package main

import "fmt"

type Student3 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main32() {
	//指针有合法指向后，才可以操作成员
	//先定义一个普通结构体变量
	var s Student3
	//定义一个指针变量，保存s的地址
	var p1 *Student3
	p1 = &s
	//通过指针操作结构体成员 p1.id与(*p1).id完全等价，只能使用点运算符
	p1.id = 2
	(*p1).addr = "fujian"
	p1.age = 19
	fmt.Println("p1 = ", p1)
	//2、通过new申请一个结构体
	p2 := new(Student3)
	p2.id = 2
	p2.addr = "fujian"
	p2.age = 19
	fmt.Println("p2 = ", p2)

}
