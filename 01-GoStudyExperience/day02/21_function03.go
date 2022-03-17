package main

import "fmt"

func function08(temp ...int) {
	for _, data := range temp {
		fmt.Println("data = ", data)
	}
	fmt.Println("")
	function09(temp...) //全部元素传递给function09
	fmt.Println("")
	function09(temp[:2]...) //从temp[0]~temp[2](不包括数字temp[2]）传递过去
	fmt.Println("")
	function09(temp[2:]...) //从temp[2]开始 包括本身把后面所有的元素传递过去
}

func function09(temp ...int) {
	for _, data := range temp {
		fmt.Println("data = ", data)
	}
}

func main2103() {
	function08(5, 3, 9)
}
