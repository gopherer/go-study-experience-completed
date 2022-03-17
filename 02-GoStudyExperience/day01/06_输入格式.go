package main

import "fmt"

func main06() {
	fmt.Println("hello")
	fmt.Print("go")

	a := 10
	b := 3.145
	//%d %f 占位符
	//%f 默认保留小数点后6位
	fmt.Printf("a = %d",a)
	fmt.Printf("b = %f",b)
	fmt.Printf("b = %.2f",b)//%.2f保留小数点后两位，第三位四舍五入

}
