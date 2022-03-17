package main

type Emp struct {
	Id   int
	Name string
	Next *Emp
}
type EmpLink struct {
	Head *Emp
}
type HashTable struct {
	LinkArr [7]EmpLink
}

func main() {

}
