package main

import "fmt"

func main07() {
	a := 20
	b := 10

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) //除数不能为0
	fmt.Println(a % b) //b 不能等于0
}

func main0701() { //自增 自减 不能运用于表达式 避免二义性 a=b++//err
	a := 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	b := 3.15
	b++
	fmt.Println(b)
	b--
	fmt.Println(b)
}
