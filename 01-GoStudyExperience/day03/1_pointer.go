package main

import "fmt"

func main01() {
	var a int = 10
	//每个变量有两层含义：变量的内存，变量的地址
	fmt.Printf("a = %d\n", a) //变量的内存
	fmt.Printf("&a = %p\n", &a)
	fmt.Printf("&a = %v\n", &a) //%v自动匹配类型
	//保存某个变量的地址，需要指针变量类型 *int 保存int的地址 **int保存*int的地址
	//声明（定义），定义只是特殊的声明
	//定义一个变量p 类型为*int
	var p *int
	p = &a //指针变量指向谁，就是把谁的地址赋值给指针变量
	fmt.Printf("p = %v\n,&a = %v\n", p, &a)

	*p = 666
	fmt.Printf("p = %v\n", p) //*p操作的不是p的内存，是p所指向内存(a)
	fmt.Println("&a = ", &a)
}
