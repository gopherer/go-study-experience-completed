package main

import "fmt"

func main02() {
	var strarr []string = []string{"小明", "小红", "小白", "小黑"}
	var str string
	fmt.Println("请输入一个名字 ： ")
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//sort1(strarr, str)
	sort2(strarr, str)
}
func sort1(strarr []string, str string) {
	//第一种排序
	for i := 0; i < len(strarr); i++ {
		if str == strarr[i] {
			fmt.Printf("恭喜你 找到了 %v  下标为%v\n", strarr[i], i)
		} else if i == len(strarr)-1 {
			fmt.Println("很遗憾 未找到")
		}
	}
}
func sort2(strarr []string, str string) {
	//第二种排序
	var index int = -1
	for i, _ := range strarr {
		if str == strarr[i] {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("恭喜你 找到了 %v 下标为 %v\n", strarr[index], index)
	} else {
		fmt.Println("很遗憾 未找到")
	}
}
