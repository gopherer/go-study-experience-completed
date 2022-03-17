package main

import (
	"fmt"
	"strings"
)

func main01() {
	var s string
	//多行字符串 原样输出 不会转义
	s = `你好吗\我很好\你不好\好好hao`
	fmt.Println(s)
	//字符串的拼接
	s1 := "hello"
	s2 := "world"
	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s%s", s1, s2)
	s4 := fmt.Sprintln(s1, s2)
	fmt.Println(s3)
	fmt.Println(s4)
	//字符串前缀 后缀
	fmt.Println(strings.HasPrefix(s, "你好"))
	fmt.Println(strings.HasSuffix(s, "hahaha"))
	//字符出现的位置
	fmt.Println(strings.Index(s, "h"))
	fmt.Println(strings.LastIndex(s, "好"))

	str := "hello 拜拜"
	for i, v := range str {
		fmt.Printf("%d %c \n", i, v)
	}
	for i := range str {
		fmt.Println(i)
	}

}
