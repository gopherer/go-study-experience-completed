package main

import "fmt"

func testaaa() {
	fmt.Println("aaaaaaaaaaaaaa")
}
func testbbb(x int) {
	//设置recover
	defer func() {
		//recover()
		//可以打印panic的错误信息
		//fmt.Println(recover())
		if err := recover(); err != nil { //产生一个panic异常
			fmt.Println(err)
		}
	}() //别忘了()，调用此匿名函数
	var a [10]int
	a[x] = 111
}
func testccc() {
	fmt.Println("ccccccccccc")
}
func main05() {
	testaaa()
	testbbb(20)
	testccc()
}
