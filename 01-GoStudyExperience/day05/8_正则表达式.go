package main

import (
	"fmt"
	"regexp"
)

func main08() {
	buf := "abc azc a7c 999 a9c tac"

	//1、解释规则，它会解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`a.c`)
	reg = regexp.MustCompile(`a[0-9]c`)
	reg = regexp.MustCompile(`a\dc`)
	if reg == nil { //解释失败 返回nil
		fmt.Println("regexp err")
		return
	}
	//2、根据规则提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1) //-1表示查找所有匹配的
	fmt.Println("result = ", result)
}
