package main

import "fmt"

type Student4 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main33() {
	s1 := Student4{1, "mike", 'm', 18, "fujian"}
	s2 := Student4{1, "mike", 'm', 18, "fujian"}
	s3 := Student4{2, "mike", 'm', 18, "fujian"}
	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)
	//同类型的结构体允许相互赋值
	var tep Student4
	tep = s1
	fmt.Println("tep = ", tep)
}
