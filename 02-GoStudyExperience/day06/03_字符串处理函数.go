package main

import (
	"fmt"
	"strings"
)

func main03() {
	//查找一个字符串在另外一个字符串中是否出现
	str := "hello go "
	str1 := "go"
	//Contains(被查找字符串,查早字符串)
	//一般用于模糊查找
	a := strings.Contains(str, str1)
	if a {
		fmt.Println("找到了")
	} else {
		fmt.Println("未找到")
	}
}

func main0301() {
	//字符串切片
	slice := []string{"hahaha", "hello", "go"}
	//字符串连接
	str := strings.Join(slice, "!")
	fmt.Println(str)
}

func main0302() {
	str := "hello go"
	str1 := "rr"
	//查找字符串中出现另一个字符串出现的位置，返回下标 未找到返回-1
	i := strings.Index(str, str1)
	fmt.Println(i)
}

func main0303() {
	str := "go"
	//将一个字符串重复n次
	str1 := strings.Repeat(str, 3)
	fmt.Println(str1)
}

func main0304() {
	str := "oooookkkkkkk"
	//字符串替换 用作于屏蔽敏感词汇
	//如果替换次数小于0 表示全部替换
	str1 := strings.Replace(str, "o", "a", -1)
	fmt.Println(str1)
}

func main0305() {
	str1 := "123456@qq.com"
	//将字符串按照特定的标志进行切割
	slice := strings.Split(str1, "@")
	fmt.Println(slice[0])
}

func main0306() {
	str := "========a==u==ok======="
	//去掉头尾特定内容
	str1 := strings.Trim(str, "=")
	fmt.Println(str1)
}

func main0307() {
	str := "     are u ok    "
	//去掉字符串中的空格转换成切片 一般用于统计单词个数
	slice := strings.Fields(str)
	fmt.Println(slice)
	fmt.Println(len(slice))
	for _, v := range slice {
		fmt.Printf("==%s==\n", v)
	}
}

func main0308() {
	//查找
	//bool类型=strings.Contains(被查找字符串,查找字符串)
	//int类型=strings.Index(被查找字符串,查找字符串)
	//
	//组合和分割
	//string类型=strings.Join(字符串切片,标志)
	//[]string类型=strings.Split(切割字符串,标志)
	//
	//重复和替换
	//string类型=strings.Repeat(字符串,次数)
	//string类型=strings.Replace(字符串,被替换字符串,替换字符串,次数)
	//
	//去掉内容
	//stirng内容=strings.Trim(字符串,去掉字符串)
	//[]string类型=strings.Fields(字符串)

}
