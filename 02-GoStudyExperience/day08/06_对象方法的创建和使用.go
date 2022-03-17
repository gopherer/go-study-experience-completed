package main

import "fmt"

//1、定义函数类型 定义结构体名称
//2、为已存在的数据类型起别名

type Int int

func (a *Int) add(b Int) Int {
	fmt.Println(a)
	*a = 100
	return *a + b
}
func main06() {
	//编译过程
	//1、预处理
	//2、编译 如果代码出错会提示 如果没有出错会编译为汇编语言
	//3、汇编 将汇编文件转为二进制文件
	//4、链接 将支持的库链接到程序中 变为可执行程序

	//类型别名在编译的时候会进行转换
	var a Int = 10
	var b Int = 20
	sum := a.add(b)
	fmt.Println(sum)
	fmt.Printf("sum type is %T\n", sum)

}
