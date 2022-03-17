package main

import "fmt"

func main10() {
	var t complex128
	t = 2.3 + 3.5i
	fmt.Println("t = ", t)

	//自动推导类型
	t2 := 3.5 + 4.9i
	fmt.Printf("t2 type is %T\n", t2)

	//通过内建函数，取得实部和虚部
	fmt.Println("real(t2) = ", real(t2), ",imag(t2) = ", imag(t2))
}
