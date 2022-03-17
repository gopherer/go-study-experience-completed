package main

import "fmt"

func main03() {
	var a int = 10
	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	*p = 100
	fmt.Println(a)
	fmt.Println(*p)
}

func main0301() {
	//声明一个指针变量 默认指针地址为 0x0(nil)  0-255为系统占用
	var p *int
	fmt.Printf("%p\n", p)
	fmt.Println(p)
	//new(类型)  开辟对应类型的空间
	p = new(int)
	fmt.Println(p)
	*p = 100
	fmt.Println(*p)
}

func main0302() {
	//野指针 指向无效的内存地址  程序允许出现空指针不允许出校野指针
	var p *int = *int(0xc0000a0088)
	fmt.Println(p)
}
