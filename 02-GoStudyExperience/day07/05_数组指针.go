package main

import "fmt"

func main05() {
	arr := [5]int{1, 2, 3, 4, 5}
	p := &arr

	fmt.Printf("p type is %T\n", p)
	fmt.Println(len(p), len(arr))
	(*p)[0] = 100
	p[1] = 200
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(arr)

	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
}
func main0501() {
	var arr [3]int = [3]int{1, 2, 3}
	var p *[3]int
	p = &arr
	p1 := &arr[0]

	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", p1)

	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", p1)

}

func main0502() {
	arr := [10]int{1, 9, 2, 8, 3, 7, 4, 5, 6}
	p := &arr
	test1(p)
	fmt.Println(arr)
}
func test1(p *[10]int) {
	for i := 0; i < len(p)-1; i++ {
		for j := 0; j < len(p)-1-i; j++ {
			if p[j] < p[j+1] {
				p[j], p[j+1] = p[j+1], p[j]
			}
		}
	}
}
