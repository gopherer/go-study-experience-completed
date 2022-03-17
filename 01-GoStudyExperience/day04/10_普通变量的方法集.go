package main

import "fmt"

//Person8 内容
type Person8 struct {
	name string
	sex  byte
	age  int
}

//SIF1 内容
func (p Person8) SIF1() {
	fmt.Println("SIF1")
}

//SIP1 内容
func (p *Person8) SIP1() {
	fmt.Println("SIP1")
}
func main10() {
	//结构体变量的一个指针变量，它能够调用那些方法，这些方法就是一个集合，简称方法集
	p := Person8{"mike", 'n', 19}
	p.SIP1() //先把p转换为&p后再调用
}
