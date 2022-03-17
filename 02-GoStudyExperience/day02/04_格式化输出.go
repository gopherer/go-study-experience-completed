package main

import "fmt"

func main04() {
	a := 10

	fmt.Printf("a = %b\n", a)
	fmt.Printf("a = %o\n", a)
	fmt.Printf("a = %d\n", a)
	fmt.Printf("a = %x\n", a)
	fmt.Printf("a = %X\n", a)
	fmt.Println(a) //10
}
func main0401() {
	var a int = 10
	var b int = 010
	var c int = 0x10

	//十进制打印输出
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//go 语言无法输入 定义 二进制数
}

func main0402() {
	var ch byte = 'a'
	var str string = "hello"

	fmt.Printf("ch = %c\n", ch)
	fmt.Printf("str = %s\n", str)
	fmt.Printf("ch addr = %p\n", &ch)

	var b bool = false

	fmt.Printf("b = %t\n", b)
	fmt.Printf("b %T\n", b)
	fmt.Printf("b = %v\n", b) //%v 自动匹配类型
}
