package main

import (
	"fmt"
	"regexp"
)

func main09() {
	buf := "43.99 456 sfasf 1.23 3. 55.9 3.15 5.77"

	//解释正则表达式，+匹配前一个字符1次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	//提取关键信息
	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)
	
}
