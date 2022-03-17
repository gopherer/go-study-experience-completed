package main

import "fmt"

func main04() {
	var month int
	var year int
	fmt.Scan(&year, &month)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31")
	case 4, 6, 9, 11:
		fmt.Println("30")
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	}
}

//fallthrough 可跳过case 语句默认的break
