package main

import "fmt"

func main06() {
	//从键盘输入一个英文字符串统计每个字母出现的次数
	var str string
	fmt.Scan(&str)
	m := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}
	for k, v := range m {
		fmt.Printf("%c:%d  ", k, v)
	}
}
