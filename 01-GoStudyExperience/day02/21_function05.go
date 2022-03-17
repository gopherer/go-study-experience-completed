package main

import "fmt"
//多个返回值
func function13() (int, int, int) {
	return 1, 2, 3
}
//go语言推荐写法
func function14() (a int, b int, c int) {
	a = 11
	b = 22
	c = 55
	return
}

func function15() (a , b, c int) {
	a = 11
	b = 22
	c = 55
	return
}

func main2105() {
	a, b, c := function13()
	fmt.Printf("a = %d b = %d c = %d\n", a, b, c)
	z, x, s := function14()
	fmt.Printf("z = %d x = %d s =%d\n", z, x, s)
	q,w,e := function15()
	fmt.Printf("q = %d w = %d e =%d\n", q, w, e)
}
