package main

import "fmt"

func main07() {
	var score int = 90
	var num int = 90
	switch score {
	case 90, 88, 77: //运行多个表达式
		fmt.Println("优秀")
	//case 88: 重复定义
	//	fmt.Println("优秀")
	case num: //重复定义但可以编译运行
		fmt.Println("优秀")
	}
}

func main0701() {
	switch score := 90; /* score*/ { //允许包含一个表达式以分号 隔开 不建议
	case score > 90:
		fmt.Println("优秀")

	case score > 80 && score <= 90:
		fmt.Println("良好")
	}
}

func main0702() {
	switch { //用法类似于if..else
	case 10 > 9:
		fmt.Println("10 >9")
		fallthrough //可以穿透一层switch
	case 8 > 7:
		fmt.Println("8 > 7")
	default:
		fmt.Println("......")
	}
}

func main0703() {
	//类型断言
	var i interface{}
	i = "hello"

	switch t := i.(type) { //t的值为 hello
	case int:
		fmt.Println("i 的类型为int", t)
	case string:
		fmt.Println("i 的类型为string", t)
	}

}
