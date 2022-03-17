package main

import (
	"fmt"
	"strconv"
)

func main04() {
	str := "hello go"
	//将字符串转换成切片
	slice := []byte(str)

	fmt.Println(slice)
	slice[1] = 'a'
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%c", slice[i])
	}
}

func main0401() {
	//将其他类型转化为字符串
	a := true
	str := strconv.FormatBool(a)
	fmt.Println(str)

	str1 := strconv.FormatInt(123, 10) //在计算机中可以表示2-36进制
	fmt.Println(str1)

	str2 := strconv.FormatFloat(3.15454, 'f', 2, 64)
	fmt.Println(str2)

	str3 := strconv.Itoa(125)
	fmt.Println(str3)
}

func main0402() {
	//将字符串转换位其他类型
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println(b)
	}
	v, _ := strconv.ParseFloat("3.156", 64)
	fmt.Println(v)

	a, _ := strconv.ParseInt("22222", 8, 64)
	fmt.Println(a)
}

func main0403() {
	slice := make([]byte, 0, 1024)
	//将其他类型转换为字符串添加到字符切片中
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendFloat(slice, 3.44555, 'f', 2, 64)
	slice = strconv.AppendInt(slice, 33993, 16)
	slice = strconv.AppendQuote(slice, "hello go")

	fmt.Println(string(slice))
}
