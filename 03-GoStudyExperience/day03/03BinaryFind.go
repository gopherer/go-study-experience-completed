package main

import "fmt"

func main03() {
	var arr [6]int = [6]int{1, 3, 5, 7, 9, 11}
	//二分查找
	//你要查找的数
	var ValueFind int
	fmt.Println("请输入要查找的数字： ")
	_, err := fmt.Scanln(&ValueFind)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	BinaryFind(&arr, 0, len(arr)-1, ValueFind)
}
func BinaryFind(arr *[6]int, LeftIndex int, RightIndex int, Value int) {
	//先找到中间的下标
	middle := (LeftIndex + RightIndex) / 2
	//程序退出条件
	if LeftIndex > RightIndex {
		fmt.Println("很遗憾 未找到该数字")
		return
	}
	//查找
	if arr[middle] > Value { //LeftIndex --- middle-1
		BinaryFind(arr, LeftIndex, middle-1, Value)
	} else if arr[middle] < Value { //middle+1 ---- RightIndex
		BinaryFind(arr, middle+1, RightIndex, Value)
	} else {
		fmt.Println("恭喜你 找到了 下标为 ", middle)
	}

}
