package main

import "fmt"

//结构体嵌套结构体
//父类
type Person struct {
	name string
	sex  string
	age  int
}

//子类
type Student struct {
	//将Person结构体作为Student结构体成员
	//p Person
	Person //匿名字段
	class  int
	score  int
}

func main01() {
	var stu Student
	stu.Person.name = "关羽"
	stu.Person.age = 38
	stu.sex = "男"
	stu.class = 302
	stu.score = 99
	fmt.Println(stu)
}

func main0101() {
	var stu Student = Student{
		Person: Person{"张飞", "男", 36},
		class:  302,
		score:  60,
	}
	fmt.Printf("name:%s\n", stu.name)
	fmt.Printf("sex:%s\n", stu.sex)
	fmt.Printf("age:%d\n", stu.age)
	fmt.Printf("class:%d\n", stu.class)
	fmt.Printf("score:%d\n", stu.score)

}
