package main

import "fmt"

func main02() {
	var a byte = 'a'
	//byte 是uint8的别名
	fmt.Println(a) //97
	fmt.Println(97)
	fmt.Printf("a = %c\n", a)
	fmt.Printf("a = %c\n", a-32)
	fmt.Printf("a = %T\n", a) //uint8

	var b byte = '0'

	fmt.Println(b) //48
	fmt.Printf("b = %c\n", b)

	var c byte = '\n'

	fmt.Println(c) //10
	fmt.Printf("c = %c\n", c)

	c = '\t'
	fmt.Println(c) //'\t' 水平制表符，等于8个空格
	fmt.Printf("c = %c\n", c)

}
