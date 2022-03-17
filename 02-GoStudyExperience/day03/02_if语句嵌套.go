package main

import "fmt"

func main02() {
	var score int

	fmt.Scan(&score)

	if score > 700 {
		fmt.Println("我要上清华:")
		if score > 720 {
			fmt.Println("我要学计算机")
		}
	} else if score > 680 {
		fmt.Println("我要上北大")
	}

}
