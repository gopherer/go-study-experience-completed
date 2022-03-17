package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

func main08() {
	var per Person = Person{
		name: "李四",
		age:  18,
		sex:  'm',
	}
	p := &per
	fmt.Printf("p type is %T\n", p)
	(*p).name = "张三"
	p.age = 100

	fmt.Println(per)
	fmt.Println(p)
}

func test2(p *Person) {
	p.name = "冲啊"
}

func main0801() {
	per := Person{
		name: "张三",
		age:  20,
		sex:  'f',
	}
	test2(&per)
	fmt.Println(per)
}

func main0802() {
	arr := [3]Person{
		{
			name: "张三",
			age:  18,
			sex:  'f',
		},
		{
			name: "王五",
			age:  20,
			sex:  'm',
		},
		{
			name: "李六",
			age:  30,
			sex:  'f',
		},
	}

	p := &arr

	fmt.Printf("p type is %T\n", p)
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
}

func main0803() {
	m := make(map[int]*[3]Person)

	m[1] = new([3]Person)
	m[1] = &[3]Person{
		{
			name: "张三",
			age:  18,
			sex:  'f',
		},
		{
			name: "王五",
			age:  20,
			sex:  'm',
		},
		{
			name: "李六",
			age:  30,
			sex:  'f',
		},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, *v)
	}
}
