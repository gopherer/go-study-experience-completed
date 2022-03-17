package main

import "fmt"

//Person 内容
type Person struct {
	name string
	sex  byte
	age  int
}

//Student 内容
type Student struct {
	Person //只有类型，没有名字，继承了Person的成员
	id   int
	addr string
}

func main01() {
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "fujian"}
	fmt.Println("s1 = ", s1)
	//自动推导类型
	s2 := Student{Person{"mike", 'm', 19}, 2, "bj"}
	fmt.Println("s2 = ", s2)
	//%+v,显示更详细
	fmt.Printf("s2 = %+v\n", s2)
	//指定成员初始化，没有初始化的成员自动赋值为0
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v", s3)

	s4 := Student{Person: Person{name: "mike"}, id: 4}
	fmt.Printf("s4 = %+v\n", s4)

	//s5 := Student{"mike",'w'19,2,"fujian"} errors
}
