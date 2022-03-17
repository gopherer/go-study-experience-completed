package main

import "fmt"

func main10() {
	arr := []Student{
		{"张三", 10, 20},
		{"李四", 80, 30},
		{"王五", 40, 50},
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	arr = append(arr, Student{"小明", 90, 90})
	//数组 -- 值传递
	//切片 -- 地址传递 引用传递
	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []Student) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j].age > arr[j+1].age {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
