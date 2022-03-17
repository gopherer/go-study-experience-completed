package main

import "fmt"

//Per1 人——内容
type Per1 struct {
	name string
	sex  int
	age  int
}

//PrintInfo 方法
func (tem *Per1) PrintInfo() {
	fmt.Printf("name = %s sex = %c age = %d\n", tem.name, tem.sex, tem.age)
}
//PrintInfo 方法
func (tem *Stu1) PrintInfo() {
	fmt.Printf("name = %s sex = %c age = %d\n", tem.name, tem.sex, tem.age)
}

//Stu1 继承Per字段，成员和方法都继承了s
type Stu1 struct {
	Per1  //匿名字段
	id   int
	addr string
}

func main12(){
	s := Stu1{Per1{"mike", 'm', 17}, 1, "fujian"}
	//就近原则：先找本作用域的方法，找不到再用继承的方法
	s.PrintInfo()//结论：使用Stu的方法
	//显示调用继承的方法
	s.Per1.PrintInfo()
}