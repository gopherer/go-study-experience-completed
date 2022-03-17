package main

import "fmt"

type person3 struct {
	id   int
	name string
	age  int
	sex  string
}

type student5 struct {
	person3
	score int
	class int
}

func (p *person3) SayHello() {
	fmt.Printf("大家好，我是%s 我今年%d 我是%s生 \n", p.name, p.age, p.sex)
}
func (s *student5) SayHello() {
	fmt.Printf("大家好，我是%s 我今年%d 我是%s生 我的分数%d 我的班级%d\n", s.name, s.age, s.sex, s.score, s.class)
}
func main11() {
	stu := student5{person3{10, "孙权", 18, "男"}, 100, 2001}
	//子类对象 采用就近原则
	stu.SayHello()
	//父类对象
	//如果子类和父类有相同的方法名 就叫方法重写
	stu.person3.SayHello()
	fmt.Println(stu.SayHello)
	fmt.Println(stu.person3.SayHello)

}
