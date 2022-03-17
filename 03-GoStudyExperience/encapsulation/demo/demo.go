package demo

import "fmt"

type person struct {
	Name string
	age int
	salary float64
}

func NewPerson(name string)*person	{
	return &person{
		Name:name,
	}
}

func (p *person)SetAge(age int){
	if age < 0 || age >150{
		fmt.Println("输入年龄有误：")
		return
	}
	p.age = age
}
func (p *person)GetAge()int{
	return p.age
}

func (p *person)SetSalary(salary float64){
	if salary <= 3000 || salary >30000{
		fmt.Println("输入工资有误：")
		return
	}
	p.salary = salary
}
func (p *person)GetSalary()float64{
	return p.salary
}