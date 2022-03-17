package main

import "fmt"

type FC func(int, int)

func demo1(a int, b int) {
	fmt.Println(a + b)
}
func demo2(a int, b int) {
	fmt.Println(a - b)
}
func main06() {
	//函数的名字表示一个地址，函数在代码区的地址
	//fmt.Println(demo1)
	demo1(10, 20)
	f := demo1
	f(10, 20)
	//fmt.Println(f)
	fmt.Printf("f type is %T\n", f)

	var ff func(int, int)
	ff = f
	ff(1, 2)
	ff = demo2
	ff(1, 2)

	var fff FC
	fff = ff
	fff(2, 3)
	fmt.Printf("fff type is %T\n", fff)

}
