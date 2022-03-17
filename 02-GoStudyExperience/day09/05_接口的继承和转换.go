package main

import "fmt"

type humaner interface { //子集
	sayhi()
}
type personer interface { //超集
	humaner
	sing(string)
}
type student struct {
	name string
	sex  string
	age  int
}

func (s *student) sayhi() {
	fmt.Printf("大家好，我是%s 我是%s生 我今年%d岁\n", s.name, s.sex, s.age)
}
func (s *student) sing(sing string) {
	fmt.Println("我给大家唱首歌:", sing)
}
func main05() {
	//接口变量的定义
	var h humaner
	var stu student = student{"小明", "男", 18}
	h = &stu
	h.sayhi()

	//接口变量的定义
	var p personer
	p = &stu
	p.sayhi()
	p.sing("欢迎您")
}

func main0501() {
	var h humaner
	var p personer
	var stu student = student{"小明", "男", 18}

	//p = &stu
	//h = p ok
	//h.sayhi()

	//子集不包含超集
	//可以将超集赋值给子集 不能将子集赋值给超集

	//h = &stu
	//p = h err
	//p.sing("欢迎你")
	//p.sayhi()

}
