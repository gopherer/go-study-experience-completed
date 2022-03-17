package main

import (
	"fmt"
	"testing"
)

//TestMain函数可以在测试函数执行前执行一些操作
func TestMain(m *testing.M) {
	fmt.Println("TestMain 测试开始")
	//通过m.Run()来执行测试函数
	m.Run()
	fmt.Println("main test")
}
func TestHello(t *testing.T) {
	fmt.Println("TestHello hello")
	t.Run("添加测试函数", testGo)
	fmt.Println("hello test")
}

//如果函数名不是以Test开头，那么函数默认不执行，我们可以将其设置围一个子测试函数
func testGo(t *testing.T) {
	fmt.Println("testGO go")
	fmt.Println("go test")
}
