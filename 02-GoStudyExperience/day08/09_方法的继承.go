package main

import "fmt"

type person2 struct {
	id   int
	name string
	sex  string
}
type student4 struct {
	person2
	score int
	class int
}

func (p *person2) Print() {
	fmt.Println(*p)
}
func (s *student4) Print() {
	fmt.Println(*s)
}
func main09() {
	p := person2{10, "孙尚香", "女"}
	p.Print()
	stu := student4{person2{11, "刘备", "男"}, 99, 2001}
	//子类可以继承父类的结构体（属性） 也可以继承于父类的方法
	//父类不可以继承与子类
	stu.Print()
	stu.person2.Print()
}
