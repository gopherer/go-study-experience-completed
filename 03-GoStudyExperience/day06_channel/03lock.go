package main

import (
	"fmt"
	"sync"
)

var (
	m    = make(map[int]int, 10)
	lock sync.Mutex
)

func demo(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	m[n] = res
	lock.Unlock()
}
func main03() {

	for i := 1; i <= 20; i++ {
		go demo(i)
	}
	lock.Lock()
	for i, v := range m {
		fmt.Printf("map[%v] = %v\n", i, v)
	}
	lock.Unlock()
}
