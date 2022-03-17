package main

import "fmt"

func main09() {
	var arr [5]int = [5]int{10, 20, 40, 50, 90}
	fmt.Println("arr = ", arr)

	var slice []int = arr[1:3] //slice =  [20 40]
	fmt.Println("slice = ", slice)

	arr[1] = 10000
	fmt.Println("arr-", arr)     //arr- [10 10000 40 50 90]
	fmt.Println("slice-", slice) //slice- [10000 40]
	slice = append(slice, 30000000)
	fmt.Println("arr =", arr)      //arr = [10 10000 40 30000000 90]
	fmt.Println("slice = ", slice) //slice =  [10000 40 30000000]
	//arr[1] addr = 0xc000010308 -------- slice[0] addr = 0xc000010308
	fmt.Printf("arr[1] addr = %p -------- slice[0] addr = %p\n", &arr[1], &slice[0])
	slice = append(slice, 1000, 2000)
	fmt.Println("arr =", arr)      //arr = [10 10000 40 30000000 90]
	fmt.Println("slice = ", slice) //slice =  [10000 40 30000000 1000 2000]
	//arr[1] addr = 0xc000010308 -------- slice[0] addr = 0xc00000e2c0
	fmt.Printf("arr[1] addr = %p -------- slice[0] addr = %p\n", &arr[1], &slice[0])

	//append允许追加切片 append 会新建切片
	slice2 := append(slice, slice...)
	fmt.Println("slice2 = ", slice2)
	slice = append(slice, slice...)
	fmt.Println("slice = ", slice)

	slice3 := make([]int, 5)
	copy(slice3, slice)
	fmt.Println("slice3 = ", slice3)
	//slice =  [10000 40 30000000 1000 2000 10000 40 30000000 1000 2000]
	//slice3 =  [10000 40 30000000 1000 2000]
	//就copy到len的长度为止

	var slice4 []int = []int{10, 20}
	slice4 = slice
	fmt.Println("slice = ", slice)
	fmt.Println("slice4 = ", slice4)
	//slice =  [10000 40 30000000 1000 2000 10000 40 30000000 1000 2000]
	//slice4 =  [10000 40 30000000 1000 2000 10000 40 30000000 1000 2000]
	//不受len长度影响
}
