package main

import "fmt"

//Person4 内容
type Person4 struct {
	name string
	sex  byte
	age  int
}

//Student4 内容
type Student4 struct {
	*Person4 //指针类型
	id       int
	addr     string
}

func main05() {
	s := Student4{&Person4{"mike", 'g', 18}, 999, "sz"}
	fmt.Println(s)
	fmt.Println(s.name, s.sex, s.age, s.id, s.addr)

	//先定义变量
	var s1 Student4
	s1.Person4 = new(Person4) //分配空间
	s1.name = "yoyo"
	s1.sex = 'g'
	s1.age = 18
	s1.id = 89
	s1.addr = "fj"
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

}
