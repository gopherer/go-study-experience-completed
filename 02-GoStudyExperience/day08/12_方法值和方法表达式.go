package main

import "fmt"

type student6 struct {
	name string
	age  int
	sex  string
}

//函数类型func()
func (s *student6) PrintInfo() {
	fmt.Println(*s)
}

//函数类型func(string,int,string)
func (s *student6) EditInfo(name string, age int, sex string) {
	s.name = name
	s.sex = sex
	s.age = age
}

//函数类型 func()
func test() {
	fmt.Println("hello go")
}
func main12() {
	stu := student6{"甄姬", 18, "女"}
	f := stu.PrintInfo
	fmt.Printf("f type is %T\n", f)
	f()
	f = test
	fmt.Printf("f type is %T\n", f)
	f()
}
func main1201() {
	var s student6
	//通过方法为对象赋值
	s.EditInfo("武则天", 30, "女")
	f := s.EditInfo
	fmt.Println(s)
	fmt.Printf("f type is %T\n", f)
	f("李世民", 30, "男")
	fmt.Println(s)
	//虽然对象相同 但是函数类型不同不能赋值
	//f = s.PrintInfo //err

}
