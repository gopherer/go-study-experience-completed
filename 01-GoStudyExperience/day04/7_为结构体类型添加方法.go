package main

import "fmt"

//Person5 内容
type Person5 struct {
	name string
	sex  byte
	age  int
}

//PrintInfo 带有接收者的函数叫方法
func (tep Person5) PrintInfo() {
	fmt.Println("tep = ", tep)
}

//SetInfo 通过一个函数，给成员赋值
func (tep *Person5) SetInfo(n string, s byte, a int) {
	tep.name = n
	tep.sex = s
	tep.age = a
}

func main07() {
	//定义同时初始化
	p := Person5{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person5
	//var p2 *Person5
	//p2 = new(Person5)
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()
}

//只要接收者类型不一样 函数就算同名也算不同的方法 
