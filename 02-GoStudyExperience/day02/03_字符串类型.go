package main

import "fmt"

func main03() {
	a := 'a'
	b := "a" //由'a' '\0'组成 '\0'为字符串的结束标准
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("a = %c\n", a)
	fmt.Printf("b = %s\n", b)
}

func main0301() {
	var a string = "hello go"
	count := len(a)
	fmt.Println(count)
	//在go语言中 汉字占3个字符，与Linux下汉字所占字符大小统一
	var str string = "你好"
	count = len(str)
	fmt.Println(count)
}

func main0302() {
	str1 := "hello"
	str2 := "go"
	//字符串的相加 +
	str := str1 + str2
	fmt.Println(str)
}
