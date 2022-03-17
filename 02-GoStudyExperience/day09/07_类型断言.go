package main

import "fmt"

func main07() {
	var i interface{}

	i = 100
	//值,值的判断:=接口变量.(数据类型)
	value, ok := i.(int)
	if ok {
		fmt.Println("整型数据：", value)
	} else {
		fmt.Println("错误")
	}

}
func demo() {
	fmt.Println("demo hello go")
}
func main0701() {
	var i []interface{}

	i = append(i, 10, 2.33, "hello", demo)

	for _, k := range i {
		if value, ok := k.(int); ok {
			fmt.Println("整型数据：", value)
		} else if value, ok := k.(float64); ok {
			fmt.Println("浮点型数据:", value)
		} else if value, ok := k.(string); ok {
			fmt.Println("字符串类型:", value)
		} else if value, ok := k.(func()); ok {
			value()
		}
	}
}
