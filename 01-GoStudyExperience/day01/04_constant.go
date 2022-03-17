package main

import "fmt"

func main04() {
	//常量：程序运行期间，不可改变的量，常量声明需要const
	const a int = 10
	//a = 20//err
	fmt.Println("a = ", a)

	const b = 20
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)

	const (
		i int     = 10   //i = 10
		j float64 = 4.88 //j = 4.88
	)
	fmt.Printf("i = %d, j = %f", i, j)

}
