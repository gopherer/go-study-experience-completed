package main

import "fmt"

func main18() {
	str := "apple"
	//通过for打印每一字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c ", i, str[i])
	}
	fmt.Println("")

	//迭代打印每个元素，默认返回2个值：一个是元素的位置，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d] = %c ", i, data)
	}
	fmt.Println("")

	for i := range str {
		fmt.Printf("str[%d] = %c ", i, str[i]) //第二个返回值，默认丢弃，返回元素位置下标
	}
	fmt.Println("")

	/*for i, _ := range str {
		fmt.Printf("str[%d] = %c ", i, str[i]) //第二个返回值，默认丢弃，返回元素位置下标
	}
	fmt.Println("")*/

}
