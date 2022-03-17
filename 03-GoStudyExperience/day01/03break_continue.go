package main

import "fmt"

func main03() {
	//flag:
	//	for i := 0; i < 5; i++ {
	//		for j := 0; j < 5; j++ {
	//			if i == 2 {
	//				break flag //跳出多层循环
	//			}
	//			fmt.Println(i, j)
	//		}
	//	}
	//	fmt.Println("hell")
	//for i := 0; i < 5; i++ {
	//flag:
	//	for j := 0; j < 5; j++ {
	//		if i == 2 {
	//			break flag //跳出多层循环
	//		}
	//		fmt.Println(i, j)
	//	}
	//}
	//flag:
	//	for i := 0; i < 5; i++ {
	//		for j := 0; j < 5; j++ {
	//			if i == 2 {
	//				continue flag
	//			}
	//			fmt.Println(i, j)
	//		}
	//	}

	//for i := 0; i < 5; i++ {
	//flag:
	//	for j := 0; j < 5; j++ {
	//		if i == 2 {
	//			continue flag
	//		}
	//		fmt.Println(i, j)
	//	}
	//}
flag:
	for i := 0; i < 2; i++ {
		if i == 0 {
			fmt.Println("hell")
			continue flag
		}
		fmt.Println("hel2222222222221111111111")
	}
	fmt.Println("hell1111111111111")
}
