package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name  string
	Age   int
	Score int
}
type HeroSlice []Hero

//实现接口所需的三个方法
func (hero HeroSlice) Len() int {
	return len(hero)
}
func (hero HeroSlice) Less(i, j int) bool {
	return hero[i].Name < hero[j].Name
}
func (hero HeroSlice) Swap(i, j int) {
	hero[i], hero[j] = hero[j], hero[i]
}

func main04() {
	var heroSlice HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name:  fmt.Sprintln("英雄~", rand.Intn(100)),
			Age:   rand.Intn(80),
			Score: rand.Intn(100),
		}
		heroSlice = append(heroSlice, hero)
	}
	fmt.Println("排序前''''''")
	for _, v := range heroSlice {
		fmt.Println(v.Name, v.Age, v.Score)
	}
	sort.Sort(heroSlice)
	fmt.Println("排序后''''''")
	for _, v := range heroSlice {
		fmt.Println(v.Name, v.Age, v.Score)
	}
}
