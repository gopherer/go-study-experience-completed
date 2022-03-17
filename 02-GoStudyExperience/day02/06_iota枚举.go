package main

import "fmt"

func main06() {
	//const (
	//	a    = iota       //0
	//	b, c = iota, iota //1
	//	d    = iota       //2
	//)

	//const (
	//	a = iota //0
	//	b        //1
	//	c        //2
	//	d        //3
	//)

	const (
		a = iota //0
		b = 10   //10
		c = 20   //10
		d        //20
		e = iota //4
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
