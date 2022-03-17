package main

import "fmt"

var age int = 10 //ok
//name := "jack" //err 等价于var name string  name = "jack" ,函数外不不允许出现赋值语句

func main05() {
	fmt.Println("age = ", age)
	//fmt.Println("name = ",name)
}
