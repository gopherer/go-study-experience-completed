package main

import "fmt"

func main01() {
label:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				continue label
			}
			fmt.Println(i, j)
		}
	}
}
