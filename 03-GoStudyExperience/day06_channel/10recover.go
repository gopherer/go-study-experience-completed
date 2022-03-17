package main

import "fmt"

func hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
}
func test1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test1 err", err)
		}
	}()
	var m map[int]string
	m[0] = "golang"
}
func main10() {
	go hello()
	go test1()
	for {
	}
}
