package main

import "fmt"

func main07() {
	var arr [2][3]int

	fmt.Println("二维数组的行：", len(arr))
	fmt.Println("二维数组的列：", len(arr[0]))

	arr[1][2] = 100
	arr[0][2] = 999
	fmt.Println(arr)

	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr[0]); j++ {
	//		fmt.Println(arr[i][j])
	//	}
	//}

	for i, v := range arr {
		fmt.Println(i, v)
		fmt.Printf("v type is %T\n", v)
		for _, data := range v {
			fmt.Println(data)
		}
	}
}

func main0701() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr)
	var arr1 [2][3]int = [2][3]int{{2, 3}, {2}}
	fmt.Println(arr1)
	var arr2 [2][3]int = [2][3]int{1: {1: 1}}
	fmt.Println(arr2)
	arr3 := [2][3]int{{2, 3}, {2}}
	fmt.Println(arr3)
}
