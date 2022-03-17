package main

import (
	"fmt"
	"strconv"
)

func main07() {
	//转换为字符串追加到字节切片
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//fmt.Println("slice = ", string(slice))
	//第二个数为要追加的数，第三个数为指定以10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "hellogo")
	fmt.Println("slice = ", string(slice)) //转换为string后打印

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)

	//'f'指打印格式，以小数方式,-1 指小数点位数(紧缩模式),64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str = ", str)

	//整形转字符串，常用
	str = strconv.Itoa(666)
	fmt.Println("str = ", str)

	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}
	//把字符串转换为整形
	a, _ := strconv.Atoi("66666")
	fmt.Println("a = ", a)
}
