package main

import "fmt"

func main03(){
	a := 10
	b := 3.15
	c :="hello"
	//不同的类型在内存中开辟的空间不相同
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c+" go")
	//%T打印数据对应的类型
	fmt.Printf("a = %T\n",a)
	fmt.Printf("b = %T\n",b)
	fmt.Printf("c = %T\n",c)
}
//交换两个变量的值
func main0301(){
	a := 10
	b := 20
	//通过第三方变量交换
	c := a
	a = b
	b = c

	fmt.Println(a)
	fmt.Println(b)

	//通过运算实现交换
	a = a+b
	b = a-b
	a = a-b

	fmt.Println(a)
	fmt.Println(b)

	a,b = b,a

	fmt.Println(a)
	fmt.Println(b)
}
