package main

import "fmt"

//定义接口
type num interface {
	Result() int
}

//父类
type number struct {
	num1 int
	num2 int
}

//加法子类
type add struct {
	number
}

//减法子类
type sub struct {
	number
}

//乘法子类
type mul struct {
	number
}

//工厂类
type factory struct {
}

//加法方法
func (a *add) Result() int {
	return a.num1 + a.num2
}

//减法方法
func (s *sub) Result() int {
	return s.num1 - s.num2
}

//乘法方法
func (m *mul) Result() int {
	return m.num1 * m.num2
}

//工厂的方法
func (f *factory) Result(num1, num2 int, ch string) {
	switch ch {
	case "+":
		a := add{number{num1, num2}}
		Result(&a)
	case "-":
		var b sub
		b.num1 = num1
		b.num2 = num2
		Result(&b)
	case "*":
		m := mul{number{num1, num2}}
		Result(&m)
	}
}

//多态
func Result(n num) {
	v := n.Result()
	fmt.Println(v)
}
func main0103() {
	var f factory
	f.Result(10, 20, "*")
}
func main0102() {
	a := add{number{10, 20}}
	Result(&a)
	b := sub{number{10, 20}}
	Result(&b)
}
func main0101() {
	var n num
	var a add = add{number{10, 20}}
	var b sub = sub{number{10, 20}}
	var v int
	n = &a
	v = n.Result()
	fmt.Println(v)
	n = &b
	v = n.Result()
	fmt.Println(v)
}
func main01() {
	var a add
	a.num1 = 10
	a.num2 = 20
	v := a.Result()
	fmt.Println(v)

	s := sub{number{10, 20}}
	v = s.Result()
	fmt.Println(v)
}
