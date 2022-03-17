package main

import "fmt"

//Person7 内容
type Person7 struct {
	name string
	sex  byte
	age  int
}

//SIF 内容
func (p Person7) SIF() {
	fmt.Println("SIF")
}

//SIP 内容
func (p *Person7) SIP() {
	fmt.Println("SIP")
}
func main09() {
	//结构体变量的一个指针变量，它能够调用那些方法，这些方法就是一个集合，简称方法集
	p := &Person7{"mike", 'n', 19}
	p.SIF()
	(*p).SIF()//把(*p)转换成p后在调用，等价于上面个

	p.SIP()//内部做了转换，先把指针p,转换为*p后再调用
	(*p).SIP()
}
