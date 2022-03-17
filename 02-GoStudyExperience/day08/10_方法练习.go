package main

import "fmt"

//父类
type Human struct {
	name string
	age  int
	sex  string
}

//记者子类
type Reper struct {
	Human
	hobby string
}

//程序员子类
type Dever struct {
	Human
	worktime int
}

//父类方法
func (h *Human) SayHi() {
	fmt.Printf("大家好，我叫%s,我今年%d,我是%s生 ", h.name, h.age, h.sex)
}

//记者方法
func (r *Reper) SayHello() {
	r.SayHi()
	fmt.Printf("我的爱好是%s\n", r.hobby)
}

//程序员方法
func (d *Dever) SayHello() {
	d.SayHi()
	fmt.Printf("我的工作年限%d年\n", d.worktime)
}
func main10() {
	r := Reper{Human{"卓伟", 50, "男"}, "偷拍"}
	r.SayHello()
	d := Dever{Human{"图灵", 32, "男"}, 10}
	d.SayHello()
}
