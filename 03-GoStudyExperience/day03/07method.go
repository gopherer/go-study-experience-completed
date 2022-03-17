package main

import "fmt"

type Student struct {
	Name string
	Age  int
}
type Integer int

func (stu *Student) String() string {
	res := fmt.Sprintf("name = %s age = %d", stu.Name, stu.Age)
	return res
}
func (stu *Student) Change() string {
	str := fmt.Sprintln("hello go")
	return str
}
func (stu Student) Print() {
	stu.Name = "mary" //值传递 不会影响main函数的stu
	fmt.Println("stu.Name = ", stu.Name)
}
func (i *Integer) Change() {
	*i = 100
}

func main07() {
	var stu Student = Student{"tom", 10}
	res := stu.String()
	fmt.Println(res)
	fmt.Println(&stu)
	fmt.Println(stu.String())
	fmt.Println(stu.Change())

	var i Integer
	i = 10 //编辑器 底层会对i 优化 实际传入 (&i).Change()
	i.Change()
	fmt.Println("i = ", i)

	(&stu).Print() //实际传入的值 为接收者类型为主 底层优化为stu.Print
	fmt.Println(stu)
}
