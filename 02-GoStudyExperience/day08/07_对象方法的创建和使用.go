package main

import "fmt"

type student2 struct {
	name  string
	sex   string
	score int
	class int
}

//结构体类型可以作为对象的类型
func (s student2) Print() {
	fmt.Println("姓名 = ", s.name)
	fmt.Println("性别 = ", s.sex)
	fmt.Println("分数 = ", s.score)
	fmt.Println("班级 = ", s.class)
}
func Print() {
	fmt.Println("hello go")
}
func main07() {
	//创建对象
	stu := student2{"林黛玉", "女", 100, 2001}
	//对象.方法
	stu.Print()

	stu1 := student2{"贾宝玉", "女", 99, 2001}
	stu1.Print()

	Print()
	//打印错误日志使用函数
	//print()
	//对象和方法名可以和函数重名 但相同的对象方法名不可以相同
	//打印函数在代码区的地址
	fmt.Println(stu.Print)
	fmt.Println(stu1.Print)
	fmt.Println(Print)
}
