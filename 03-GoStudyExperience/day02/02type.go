package main

import "fmt"

type myInt int
type myFunc func(int, int) int

func add(n1 int, n2 int) (sum int) {
	sum = n1 + n2
	return
}
func main02() {

	var n1 myInt = 10
	var n2 int = 20

	fmt.Printf(" n1 type is  %T\n n2 type is  %T\n", n1, n2)

	var f myFunc = add
	f1 := add
	fmt.Printf(" f type is %T\n f1 type is %T\n add type is %T\n", f, f1, add)

	sum := f(19, 20)
	fmt.Println("sum = ", sum)

}
