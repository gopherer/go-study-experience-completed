package main

import "fmt"

func main08() {
	a := 10
	b := 20

	var f func(int, int)

	f = func(a int, b int) {
		fmt.Println(a + b)
	}

	f(a, b)
	fmt.Printf("f type is %T\n", f)
}

func main0801() {
	v := func(a int, b int) int {
		return a + b
	}
	fmt.Printf("v type is %T\n", v)

	h := v(10, 20)

	fmt.Println(h)

	vv := func(a int, b int) int {
		return a + b
	}(10, 20)
	fmt.Printf("vv type is %T\n", vv)
	fmt.Println(vv)
}

func main0802() {
	a := 10
	b := 20
	v := func() int {
		return a + b
	}

	a = 100
	b = 200
	fmt.Println(v())
}
