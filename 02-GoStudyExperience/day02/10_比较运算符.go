package main

import "fmt"

func main10() {
	a := 10
	b := 20

	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a == b)
}
func main1001() {
	a := 10
	b := 20
	c := a > b
	fmt.Printf("c type %T\n", c)
	c = a+20 > b
	fmt.Println(c)

}
