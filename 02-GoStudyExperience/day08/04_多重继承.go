package main

import "fmt"

type human struct {
	name string
	sex  string
}

type being struct {
	id    int
	class int
}
type girl struct {
	human
	being
	score int
}

func main04() {
	var stu girl
	stu.name = "貂蝉"
	stu.sex = "女"
	stu.id = 100
	stu.class = 2001
	stu.score = 100
	fmt.Println(stu)

	st := girl{human{name: "吕布", sex: "男"}, being{id: 10, class: 2001}, 100}
	fmt.Println(st)
}
