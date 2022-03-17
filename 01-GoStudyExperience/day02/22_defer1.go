package main

import "fmt"

func test(x int) {
	result := 100 / x //x为0时，产生异常
	fmt.Println("result = ", result)
}
func main2201() {
	defer fmt.Println("11111111111")
	defer fmt.Println("2222222222222222")

	//调用一个函数 导致内存出现问题
	//test(0)
	defer test(0)

	defer fmt.Println("3333333333")
}

//函数中含有多个defer语句，它们会以LIFO(先进后出)的顺序执行，
//哪怕函数或某个defer延迟发生错误，这些调用依旧会被执行------------至下而上运行（个人理解）
