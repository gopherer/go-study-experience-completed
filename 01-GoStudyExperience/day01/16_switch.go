package main

import "fmt"

func main16() {
	var num int
	fmt.Scan(&num)

	switch num { //switch 后面接变量本身
	case 1:
		fmt.Println("hello")
		break //go语言保留了break关键字，跳出switch语句，不写默认包含
	case 2:
		fmt.Println("world")
		fallthrough //不跳出switch语句，后面的无条件执行
	case 3:
		fmt.Println("你瞅啥")
	default:
		fmt.Println("瞅你乍地")
	}
}
