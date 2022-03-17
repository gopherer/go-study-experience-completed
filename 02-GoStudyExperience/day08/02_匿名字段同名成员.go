package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}
type student struct {
	person
	name  string
	score int
}

func main02() {
	var stu student
	stu.score = 100
	stu.age = 19
	stu.sex = "男"
	//给子类赋值
	stu.name = "李白"
	//给父类赋值
	stu.person.name = "白白"
	fmt.Println(stu)

	st := student{
		person: person{"李四", 18, "女"},
		name:   "李青",
		score:  99,
	}
	fmt.Println(st)
}
