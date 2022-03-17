package main

import "C"
import (
	"os"
)

type User struct {
	Number1 string
	Number2 string
}

var user User

func main() {
	result := UI()
	if result == 1 {
		result := user.Multiplication()
		user.MulResultUI(result)
	} else if result == -1 {
		os.Exit(0)
	}
}
func init() {
	PreUI()
}
