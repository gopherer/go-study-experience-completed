package main

import "fmt"

type student3 struct {
	name  string
	sex   string
	score int
	class int
}

//对象不同方法名相同 不会冲突
//在方法调用中 方法的接收者为指针类型 影响
//指针类型和普通类型表示是相同的类型
func (s *student3) Print() {
	s.name = "小乔"
	fmt.Println(*s)
}
func main08() {
	stu := student3{"大乔", "女", 100, 2001}
	//地址传递（*student3） 值传递（student3）
	//一般建议选择地址传递
	stu.Print()
	fmt.Println(stu)
}
