package main

import "fmt"

func main07() {
	//声明变量
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)
	//自动推导类型
	f2 := 2.23
	fmt.Printf("f2 taye is %T\n", f2) //float64
	//float64存储小数比float32更准确
}
