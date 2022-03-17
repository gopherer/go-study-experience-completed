package main

import (
	"fmt"
	"runtime"
)

func main02() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	fmt.Println("本机CPU数:= ", num)
}
