package main

import "fmt"

type person1 struct {
	name string
	age  int
	sex  string
}
type student1 struct {
	*person1
	score int
}

func main03() {
	var stu student1
	//stu.person1 是一个指针类型 默认值为nil
	stu.score = 99
	stu.person1 = new(person1)
	stu.name = "韩星"
	stu.age = 19
	stu.sex = "女"
	fmt.Println(stu)

	fmt.Printf("name:%s\n", stu.name)
	fmt.Printf("sex:%s\n", stu.sex)
	fmt.Printf("age:%d\n", stu.age)
	fmt.Printf("score:%d\n", stu.score)

	st := student1{&person1{"小六子", 18, "男"}, 100}
	fmt.Println(st)
	fmt.Printf("name:%s\n", st.name)
	fmt.Printf("sex:%s\n", st.sex)
	fmt.Printf("age:%d\n", st.age)
	fmt.Printf("score:%d\n", st.score)
}
