package main

import "fmt"

//Humaner 接口
type Humaner interface {
	sayhi()
}
type stud struct {
	name string
	id   int
}

//stud 实现此方法
func (tem *stud) sayhi() {
	fmt.Printf("stud[%s %d] sayhi\n", tem.name, tem.id)
}

type teacher struct {
	addr  string
	group string
}

//teacher实现此方法

func (tem *teacher) sayhi() {
	fmt.Printf("teacher[%s %s] sayhi\n", tem.addr, tem.group)
}

//MyStr string
type MyStr string

//MyStr 实现此方法
func (tem *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi", *tem)
}

//WhoSayHi 多态
func WhoSayHi(i Humaner) {
	i.sayhi()
}
func main1402() {
	s := &stud{"mike", 666}
	t := &teacher{"bj", "go"}
	var str MyStr = "hello mike"
	//调用同一个函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)
	fmt.Println("\n=============")

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, i := range x {
		i.sayhi()
	}

}
func main1401() {
	//定义接口类型的变量
	var i Humaner
	//只是实现了此接口方法的类型，那么这个类型的变量(接收者类型)，就可以给i赋值
	s := &stud{"mike", 666}
	i = s
	i.sayhi()

	t := &teacher{"bj", "go"}
	i = t
	i.sayhi()

	var str MyStr = "hello mike"
	i = &str
	i.sayhi()
}
