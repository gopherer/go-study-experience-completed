package main

import "fmt"

func main1601() {
	//支持一个初始化语句，初始化语句和变量本身用分号分隔
	switch num := 4; num { //switch 后面接变量本身
	case 1:
		fmt.Println("hello")
		break //go语言保留了break关键字，跳出switch语句，不写默认包含
	case 2:
		fmt.Println("world")
		fallthrough //不跳出switch语句，后面的无条件执行
	case 3, 4, 5:
		fmt.Println("你瞅啥")
	default:
		fmt.Println("瞅你乍地")
	}

	score := 88
	switch { //可以没有条件
	case score > 90:
		fmt.Println("流弊")
	case score > 80:
		fmt.Println("还行")
	case score > 70:
		fmt.Println("咦")
	default:
		fmt.Println("差")
	}
}
