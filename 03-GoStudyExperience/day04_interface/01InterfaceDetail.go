package main

import "fmt"

type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}
type Integer int
type Stu struct {
	Name string
}
type Monster struct {
}

func (m Monster) Say() {
	fmt.Println("Monster Say()")
}
func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}
func (i Integer) Say() {
	fmt.Println("Integer Say() i = ", i)
}
func (s Stu) Say() {
	fmt.Println("hello go")
}
func main() {
	var stu Stu //结构体变量 实现Say() 实现了AInterface
	var a AInterface = stu
	a.Say()

	//自定义类型 即可 不限于struct
	var i Integer = 10
	var a1 AInterface = i
	a1.Say()

	var monster Monster //monster 实现了接口AInterface 和 BInterface 的方法
	var b BInterface
	a = monster
	a.Say()
	b = monster
	b.Hello()

}
