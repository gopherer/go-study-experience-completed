package main

import "fmt"

func BubbleSort(arr *[5]int) {
	//func len(v Type) int
	//内建函数len返回 v 的长度，这取决于具体类型：
	//数组：v中元素的数量
	//数组指针：*v中元素的数量（v为nil时panic）
	//切片、映射：v中元素的数量；若v为nil，len(v)即为零
	//字符串：v中字节的数量
	//通道：通道缓存中队列（未读取）元素的数量；若v为 nil，len(v)即为零
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	/*
		//第二次排序
		for j := 0; j < len(arr)-2; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		.....*/
}
func main01() {
	var arr [5]int = [5]int{24, 89, 57, 10, 34}
	fmt.Println("排序前 arr = ", arr)
	BubbleSort(&arr)
	fmt.Println("排序后 arr = ", arr)
}
