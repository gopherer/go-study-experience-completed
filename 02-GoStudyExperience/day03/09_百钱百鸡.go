package main

import "fmt"

func main09() {
	count := 0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			for chicken := 0; chicken <= 100; chicken += 3 {
				count++
				if cock+hen+chicken == 100 && 5*cock+3*hen+chicken/3 == 100 {
					fmt.Printf("cock:%d hen:%d chicken:%d\n", cock, hen, chicken)
				}
			}
		}
	}
	fmt.Println(count)
}
func main0901() {

	count := 0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			count++
			chicken := 100 - cock - hen
			if chicken%3 == 0 && 5*cock+3*hen+chicken/3 == 100 {
				fmt.Printf("cock:%d hen:%d chicken:%d\n", cock, hen, chicken)
			}
		}
	}
	fmt.Println(count)
}
