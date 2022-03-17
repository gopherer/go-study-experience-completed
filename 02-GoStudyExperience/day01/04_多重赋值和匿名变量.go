package main

import "fmt"

func main04() {
	//变量和值要一一对应
	a,b,c := 10,2.344,"hello"
	fmt.Println(a,b,c)
}

func main0401(){
	var a int = 10
	var b int = 20

	//var a float64 = 3.34 同一作用域下变量不能同名
	//自动推导类型时，如果有新变量和已定义变量同时存在，已定义的变量由定义转变为赋值
	a,b,c,d := 112,120,4.5,"hello"

	fmt.Println(a,b,c,d)
}

func main0402(){
	var a int = 10
	var b int = 20

	//_表示匿名变量，不接受数据
	_,b,c,d := 110,220,"hello","go"

	fmt.Println(a,b,c,d)

}


