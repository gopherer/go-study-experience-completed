package main

import (
	"03_gocode/encapsulation/demo"
	"fmt"
)

func main() {
	var p = demo.NewPerson("jock")
	p.SetAge(18)
	p.SetSalary(5000)

	fmt.Printf("p.name = %v p.age = %v p.salary = %v\n",p.Name,p.GetAge(),p.GetSalary())
}
