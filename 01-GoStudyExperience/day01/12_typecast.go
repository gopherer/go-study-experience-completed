package main

import "fmt"

func main12() {

	//这种不能转换的类型，叫不兼容类型
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	//bool类型不能转换为int类型
	//fmt.Printf("flag = %d\n", int(flag))

	//0就是假，非0为真
	//整形也不能转换为bool
	//flag = bool(1)

	var ch byte
	ch = 'a' //字符类型本质上就是整形
	var t int
	t = int(ch)
	fmt.Println("t = ", t)
}
