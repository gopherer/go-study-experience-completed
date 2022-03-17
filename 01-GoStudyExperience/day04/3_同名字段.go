package main

import "fmt"

//Person2 内容
type Person2 struct {
	name string
	sex  byte
	age  int
}

//Student2 内容
type Student2 struct {
	Person2
	id   int
	addr string
	name string
}

func main03() {
	//声明(定义一个变量)
	var s Student2

	//默认规则(就近原则),如果在本作用域找到此成员，就操作此成员
	//					如果没有找到就找继承字段的
	s.name = "mike" //操作是Student的name，还是Person的name,结论是Studnet下的
	s.sex = 'm'
	s.age = 19
	s.addr = "fujian"
	fmt.Printf("s = %+v\n", s)

	//显示调用
	s.Person2.name = "youyou" //Person下的name
	fmt.Printf("s = %+v", s)
}
