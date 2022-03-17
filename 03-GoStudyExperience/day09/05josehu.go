package main

import "fmt"

type Boy struct {
	no   int
	next *Boy
}

func addBoy(num int) *Boy {
	first := &Boy{}
	current := first

	if num < 1 && num > 0 {
		fmt.Println("输入有误")
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		if i == 1 {
			first = boy
			current = boy
			current.next = first
		} else {
			current.next = boy
			current = boy
			current.next = first
		}
	}
	return first
}
func showBoy(first *Boy) {
	if first.next == nil {
		fmt.Println("link is empty")
		return
	}
	curBoy := first
	for {
		fmt.Printf("current boy ID = %d\n", curBoy.no)
		if curBoy.next == first {
			fmt.Println("over")
			break
		}
		curBoy = curBoy.next
	}
}
func playGame(first *Boy, startNo int, count int) {
	if first == nil {
		fmt.Println("link is empty")
		return
	}
	tail := first
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	for {
		for i := 1; i <= count-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("element %d\n", first.no)
		first = first.next
		tail.next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("element %d\n", first.no)
}
func main() {
	first := addBoy(50)
	showBoy(first)
	playGame(first, 50, 3)
}
