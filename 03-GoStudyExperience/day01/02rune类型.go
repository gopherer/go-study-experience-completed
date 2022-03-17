package main

import "fmt"

func main02() {
	//rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。
	s := []rune("大白兔")
	s[0] = '小'
	fmt.Println(string(s))

	s1 := "小"
	s2 := '小'
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", s2)

	c1 := "h"
	c2 := 'h'
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)

	str := []rune("ha丹巴路")
	fmt.Println(len(str))
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	str1 := string("ha丹巴路")
	fmt.Println(len(str1))
	fmt.Println(str1)
	for i := 0; i < len(str1); i++ {
		fmt.Println(str1[i])
	}

}
