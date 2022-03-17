package main

import "fmt"

func main10() {
	var i int = 0
	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}
func main1001() {
	var sum int = 0
	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
func main1002() {
	//	fmt.Println("hello go1")
	//	fmt.Println("hello go2")
	//	goto FLAG
	//	fmt.Println("hello go3")
	//	fmt.Println("hello go4")
	//	fmt.Println("hello go5")
	//FLAG:
	//	fmt.Println("hello go6")
	//	fmt.Println("hello go7")

aaa:
	fmt.Println("hello go")
	goto aaa
	fmt.Println("yeees")

}
