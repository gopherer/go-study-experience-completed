package main

import "fmt"

//Human 接口
type Human interface { //子集
	sayhi()
}

//Pers 接口
type Pers interface { //超集
	Human //匿名字段，继承了sayhi()
	sing(lrc string)
}
type stud1 struct {
	name string
	id   int
}

//stud1 实现了sayhi()
func (tep *stud1) sayhi() {
	fmt.Printf("stud1 [%s %d]\n", tep.name, tep.id)
}

//sing 接口
func (tep *stud1) sing(lrc string) {
	fmt.Println("stud1 ", lrc)
}
func main15() {
	//定义一个接口类型的变量
	var i Pers
	s := &stud1{"mike", 888}
	i = s
	i.sayhi() //继承过来的放噶
	i.sing("哈哈啊")
}
func main1501() {
	//超集可以转换为子集，反过来不可以
	var p Pers //超集
	p = &stud1{"mike", 888}
	var i Human //子集
	//p = i error
	i = p
	i.sayhi()
}
