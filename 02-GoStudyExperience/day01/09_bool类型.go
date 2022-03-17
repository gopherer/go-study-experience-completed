package main

import "fmt"

func main09() {
	var a bool

	fmt.Println("a = ",a)//未初始化 默认值位false

	var b bool

	b = true
	fmt.Printf("b =  %v\n",b)
	fmt.Printf("b 的类型 %T\n",b)

}
