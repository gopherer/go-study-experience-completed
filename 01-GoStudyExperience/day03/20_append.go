package main

import "fmt"

func main20() {
	//如果超出原来的容量，通常情况下一2倍的容量扩容
	s := make([]int, 0, 1)
	oldcap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newcap := cap(s); oldcap < newcap {
			fmt.Printf("cap %d =======> %d\n", oldcap, newcap)
			oldcap = newcap
		}
	}
}
