package main

import "fmt"

//Per 人——内容
type Per struct {
	name string
	sex  int
	age  int
}

//PrintInfo 方法
func (tem *Per) PrintInfo() {
	fmt.Printf("name = %s sex = %c age = %d\n", tem.name, tem.sex, tem.age)
}

//Stu 继承Per字段，成员和方法都继承了
type Stu struct {
	Per  //匿名字段
	id   int
	addr string
}

func main11() {
	s := Stu{Per{"mike", 'm', 17}, 1, "fujian"}
	s.PrintInfo()
}
