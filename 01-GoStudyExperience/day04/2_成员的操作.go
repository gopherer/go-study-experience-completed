package main

import "fmt"

//Person1 内容
type Person1 struct {
	name string
	sex  byte
	age  int
}

//Student1 内容
type Student1 struct {
	Person1
	id   int
	addr string
}

func main02() {
	s := Student1{Person1{"mike", 'm', 19}, 1, "fujian"}
	fmt.Println("s = ", s)
	s.name = "yoyo"
	s.id = 8
	s.Person1 = Person1{"go", 'm', 19}
	fmt.Println(s.name, s.id, s.Person1)

}
