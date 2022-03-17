package main

import "fmt"

type mystr string

//Person3 内容
type Person3 struct {
	name string
	sex  byte
	age  int
}

//Student3 内容
type Student3 struct {
	Person3 //结构体匿名字段
	int     //基础类型匿名字段
	mystr
}

func main04() {
	s := Student3{Person3{"mike", 'm', 19}, 19, "bj"}
	fmt.Printf("s = %+v\n", s)

	s.Person3 = Person3{"go", 'm', 99}

	fmt.Println(s.name, s.age, s.int)
	fmt.Println(s.Person3)
}
