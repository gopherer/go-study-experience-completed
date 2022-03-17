package main

import "fmt"

func main05() {
	const A int = 10

	fmt.Println("A = ", A) //一般来收定义常量使用大写字母
	fmt.Printf("A = %d\n", A)

	//A = 12//err
}
func main0501() {
	const B = "hello go"
	fmt.Println("B = ", B)
	fmt.Printf("B  type %T\n", B)
	//在go语言中 常量的地址不允许访问
	//fmt.Printf("B addr = %p\n", &B)
	//var p *string
	//p = &B//err
	//*p = "gogogo"
	//fmt.Println("*p = ", *p)
}

//计算圆的面积和周长
func main0502() {
	const PI = 3.14
	var r float64

	fmt.Scan(&r)
	s := PI * r * r
	l := 2 * PI * r

	fmt.Printf("s = %.2f\n", s)
	fmt.Printf("l = %.2f\n", l)
}

func main0503() {
	//数据区 -》常量区
	//字面常量
	fmt.Println("hahaha")
	fmt.Println(3 + 5)

}
