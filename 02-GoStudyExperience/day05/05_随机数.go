package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main05() {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rand.Intn(10))
	}
}
