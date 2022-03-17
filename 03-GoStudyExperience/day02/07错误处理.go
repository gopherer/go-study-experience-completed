package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer + recover来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
		//可以处理异常
		fmt.Println("发送邮件给admin")
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}

//读取配置文件init.conf
//成功继续执行 不成功返回自定义错误
func readconf(name string) (err error) {
	if name == "init.conf" {
		return nil
	} else {
		err = errors.New("文件读取失败")
		return err //return errors.New(" .... ")
	}
}
func test02() {
	err := readconf("init1.conf")
	if err != nil {
		panic(err) //终止函数执行
	} else {
		fmt.Println("执行test02的代码")
	}

}
func main07() {
	//test()
	test02()
	fmt.Println("main代码。。。。")
}
