package main

import (
	"fmt"
	"strings"
)

//判断输入的文件是否由后缀 如果没有则加上 有则原样输出
//闭包 会保留上一次引用的值 所以传入一次可以反复使用 -- 内存逃逸
//也可以同函数实现相同的功能 推荐使用闭包
func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main04() {
	f := makeSuffix(".jpg")
	fmt.Println("处理后", f("hello"))
	fmt.Println("处理后", f("go.jpg"))
	fmt.Println("处理后", f("haha.avi"))
}
