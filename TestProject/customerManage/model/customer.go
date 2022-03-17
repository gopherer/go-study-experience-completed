package model

import (
	"fmt"
)

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  int
	Email  string
}

func NewCustomer(id int, name string, gender string, age int, phone int, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func (cus *Customer) ShowInfo() {
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", cus.Id, cus.Name, cus.Gender, cus.Age, cus.Phone, cus.Email)
}
