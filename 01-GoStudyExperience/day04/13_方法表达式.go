package main

import "fmt"

//Person9 内容
type Person9 struct {
	name string
	sex  byte
	age  int
}

//SIF1 内容
func (p Person9) SIF1() {
	fmt.Println("SIF")
}

//SIP1 内容
func (p *Person9) SIP1() {
	fmt.Println("SIP")
}

func main13() {
	p := Person9{"mike", 'f', 18}
	//方法值 f := p.SIF //隐藏了接收者
	//方法表达式
	f := (Person9).SIF1
	f(p) //显示把接收者传递过去 ======> p.SIF1()

	f2 := (*Person9).SIP1
	f2(&p) //显示把接收者传递过去 ======> p.SIP1()
}
