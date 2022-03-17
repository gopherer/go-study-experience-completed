package main

import "fmt"

type Student struct {
	name  string
	age   int
	score int
}

func main09() {
	var stu Student
	stu.name = "张三"
	stu.age = 20
	stu.score = 55
	fmt.Println(stu)

	var stu1 Student = Student{"李四", 29, 88}
	fmt.Println(stu1)
	//自动推导类型 指定成员赋值
	stu2 := Student{age: 20, name: "王五", score: 90}
	fmt.Println(stu2)

}

func main0901() {
	stu := Student{"小明", 20, 55}
	//结构体成员为string 需要和结构体最大的数据类型进行对齐
	fmt.Printf("%p\n", &stu)
	fmt.Printf("%p\n", &stu.name)
	fmt.Printf("%p\n", &stu.age)
	fmt.Printf("%p\n", &stu.score)
}

func main0902() {
	stu := Student{"小明", 20, 55}

	stu1 := stu
	stu1.name = "李四"
	fmt.Println(stu1)
	fmt.Println(stu)

	if stu1 == stu {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同")
	}
}
