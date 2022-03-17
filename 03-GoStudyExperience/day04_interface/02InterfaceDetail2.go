package main

import "fmt"

type A interface {
	test01()
}
type B interface {
	test02()
}
type C interface { //要先实现C接口必先实现A B 接口
	A
	B
	test03()
}

type Person struct {
}

func (p Person) test01() {
	fmt.Println("test01() ....")
}
func (p Person) test02() {
	fmt.Println("test02() ....")
}
func (p Person) test03() {
	fmt.Println("test03() ....")
}

type T interface {
}

func main02() {
	var per Person
	//var a A = per
	//var b B = per
	//a.test01()
	//b.test02()
	var c C = per
	c.test02()
	c.test01()
	c.test03()

	var t T = per
	fmt.Println(t)
	var t2 interface{} = per
	fmt.Println(t2)
	t2 = 100
	t = 200
	fmt.Println(t2, t)
}
