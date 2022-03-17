package main

import "fmt"

func main16() {
	//空接口万能类型，可保存任意类型的值
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)
}

// func xxx(arg ...interface{}){ 不定参数可接收任意类型值

// }
