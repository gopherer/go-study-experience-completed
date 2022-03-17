package main

import "fmt"

func main09() {
	var str1 string
	str1 = "avs"
	fmt.Println("str1 = ", str1)
	fmt.Printf("str1 = %s\n", str1)

	//自动推导类型
	str2 := "apple" //字符串都隐藏了一个结束字符'\0'
	fmt.Printf("str2 type is %T\n", str2)

	//内建函数，len()可以测字符串的长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))
	//只想操作字符串的某个字符，从0开始操作
	fmt.Printf("str2[0] = %c, str2[1] = %c\n", str2[0], str2[1])

}
