package main

import "fmt"

func main17() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	//[low:high:max] 去下表low开始的元素，len=high-low cap = max-low
	s := arr[:] //[0:len(arr):len(arr)]
	fmt.Println("s = ", s)
	fmt.Printf("len = %d cap = %d\n", len(s), cap(s))
	//操作某个元素和操作数组的方式一样
	data := arr[0]
	fmt.Println("data = ", data)
	data = arr[1]
	fmt.Println("data = ", data)

	s1 := arr[3:5:6] //a[3],a[4] len=5-3 cap=6-3
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d cap = %d\n", len(s1), cap(s1))
	s1 = append(s1, 888)
	fmt.Println("s1 = ", s1)

	s2 := arr[:6] //下标从0开始 取6个元素容量也为6 常用
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d cap = %d\n", len(s2), cap(s2))

	s3 := arr[3:] //下标从3开始取到结尾
	fmt.Println("s3 = ", s3)
	fmt.Printf("len = %d cap = %d\n", len(s3), cap(s3))

}
