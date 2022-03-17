package main

import "fmt"

type boy struct {
	name string
	sex  string
}
type teacher struct {
	boy
	class int
}
type school struct {
	teacher
	addr string
}

func main05() {
	var sch school
	sch.name = "喜羊羊"
	sch.sex = "男"
	sch.class = 2004
	sch.addr = "羊村"
	fmt.Println(sch)

	sc := school{teacher{boy{"美羊羊", "女"}, 3004}, "羊村"}
	fmt.Println(sc)
}
