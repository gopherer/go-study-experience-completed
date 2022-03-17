package main

import "fmt"

func main09() {
	a := 10
	b := 20

	p := &a
	fmt.Printf("p type is %T\n", p)
	pp := &p
	fmt.Printf("pp type is %T\n", pp)
	*pp = &b
	**pp = 100
	fmt.Println(a, b)
	fmt.Println(*p, **pp)
}
