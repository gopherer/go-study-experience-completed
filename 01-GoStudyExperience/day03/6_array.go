package main

import "fmt"

func main06() {
	var id [30]int

	for i := 1; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}
}
