package main

import (
	"fmt"
	"unsafe"
)

func main08() {
	var str string = "hello world"
	//传统写法
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	//如果字符串有汉字 会出错
	str = "hahaha北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	//解决方法
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\n", str1[i])
	}
	fmt.Println(unsafe.Sizeof(str1[0]))
	//for-range 按字节遍历
	for i, v := range str1 {
		fmt.Printf("i = %d v = %c\n ", i, v)
	}
}
