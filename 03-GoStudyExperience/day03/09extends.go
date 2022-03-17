package main

import "fmt"

type student struct {
	Name string
	Age int
	Score int
}
type Pupil struct {
	student
}
type Graduate struct {
	student
}

//type Graduate struct {
//	*student
//}
//type Graduate struct {
//	 stu student //组合
//}

func (stu *student) SetScore(score int){
	//判断....
	stu.Score = score
}
func (stu *student) ShowInfo(){
	fmt.Printf("Name = %v Age = %v Score = %v\n",stu.Name,stu.Age,stu.Score)
}
func (pupil *Pupil)testing(){
	fmt.Println("小学生正在考试.....")
}
func (graduate *Graduate)testing(){
	fmt.Println("大学生正在考试....")
}
func main09() {
	var pupil = &Pupil{student{
		Name:  "jack",
		Age:   5,
	},}
	pupil.testing()
	pupil.student.SetScore(80)
	pupil.student.ShowInfo()

	var graduate  = Graduate{student{
		Name:  "mary",
		Age:   18,
	},}
	graduate.testing()
	graduate.SetScore(89)
	graduate.ShowInfo()

	var a A
	a.int = 6
	fmt.Println(a.int)
}

type A struct {
	int
	//int err 无法区分多个int
	//i int  ok
}
