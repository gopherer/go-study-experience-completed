package main

import (
	"fmt"
	"unsafe"
)

//结构体不能嵌套本身结构体
//链表的存储格式
type student7 struct {
	*student7
	name string
	age  int
	sex  string
}

func main13() {
	stu := student7{name: "李四"}
	fmt.Println(unsafe.Sizeof(stu))
	fmt.Println(stu)
}
