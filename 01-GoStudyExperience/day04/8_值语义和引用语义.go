package main

import "fmt"

//Person6 内容
type Person6 struct {
	name string
	sex  byte
	age  int
}

//SetInfoValue 修改成员变量的值
//接收者为普通变量，非指针，值语义，值传递为一份拷贝
func (p Person6) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ", p)
	fmt.Printf(" SetInfoValue &p = %p\n", &p)
}

//SetInfoPointer 接收者为指针变量，引用传递，地址传递
func (p *Person6) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ", p)
	fmt.Printf(" SetInfoPoint p = %p\n", p)
}
func main08() {
	s := Person6{"go", 'm', 22}
	fmt.Println("s = ", s)
	fmt.Printf("&s = %p\n", &s) //打印地址

	//值语义
	s.SetInfoValue("mike", 'm', 23)
	fmt.Println("s = ", s)

	//引用语义
	(&s).SetInfoPointer("yoyo", 'f', 98)
	fmt.Println("s = ", s) //打印内容
}
