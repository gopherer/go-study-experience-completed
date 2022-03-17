package main

import (
	"fmt"
	"strings"
)

func main06() {
	//"hello",中是否包含"hello",包含返回true,不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	//Join 组合
	s := []string{"abc", "hello", "yoyo", "go"}
	buf := strings.Join(s, "!")
	fmt.Println("buf = ", buf)

	//Index 查找字符串的位置
	fmt.Println(strings.Index("abchellogo", "go"))
	fmt.Println(strings.Index("hellogogo", "abc")) //不包含字符串返回-1

	buf = strings.Repeat("go", 2)
	fmt.Println("buf = ", buf) //"gogo"

	//Trim去掉两头的字符
	buf = strings.Trim("          are u ok?           ", " ")
	fmt.Println("buf = ", buf)

	//去掉空格，把元素放入切片中
	s1 := strings.Fields("            are u ok?              ")
	//fmt.Println("s1 = ",s1)
	for i, data := range s1 {
		fmt.Println(i, " ", data)
	}
}
