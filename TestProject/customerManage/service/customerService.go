package service

import (
	"Project/customerManage/model"
	"fmt"
)

type CustomerList struct {
	customer []model.Customer
	number   int
}

//func (cus *CustomerList) newCustomerList() *CustomerList {
//	//var cusList = CustomerList{}
//	cus.number = 1
//	var customer = *model.NewCustomer(1, "jack", "男", 18, 123456, "123@qq.con")
//	//cusList.customer[cusList.number] = *model.NewCustomer(1, "jack", "男", 18, 123456, "123@qq.con") //slice 未make 直接使用下标赋值会panic
//	cus.customer = append(cus.customer, customer)
//	return cus
//}
func (cus *CustomerList) ShowList() {

	fmt.Println("----------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(cus.customer); i++ {
		cus.customer[i].ShowInfo()
	}
	fmt.Println("--------------客户列表完成--------------")
}
func (cus *CustomerList) Add() {
	var name string
	fmt.Print("姓名:")
	_, err := fmt.Scanln(&name)
	ErrHandle(err)
	var gender string
	fmt.Print("性别:")
	_, err = fmt.Scanln(&gender)
	ErrHandle(err)
	var age int
	fmt.Print("年龄:")
	_, err = fmt.Scanln(&age)
	ErrHandle(err)
	var phone int
	fmt.Print("电话:")
	_, err = fmt.Scanln(&phone)
	ErrHandle(err)
	var email string
	fmt.Print("电邮:")
	_, err = fmt.Scanln(&email)
	ErrHandle(err)
	cus.number++
	var customer = *model.NewCustomer(cus.number, name, gender, age, phone, email)
	cus.customer = append(cus.customer, customer)
}
func (cus *CustomerList) findId(id int) int {
	index := -1
	for i, v := range cus.customer {
		if v.Id == id {
			index = i
		}
	}
	return index
}
func (cus *CustomerList) delete(id int) {
	index := cus.findId(id)
	if index == -1 {
		fmt.Println("很遗憾,未找到该用户.")
		return
	}
	fmt.Print("你确定要删除吗？按y删除其他任意键退出.")
	tem := "a"
	_, err := fmt.Scanln(&tem)
	ErrHandle(err)
	if tem == "y" || tem == "Y" {
		cus.customer = append(cus.customer[0:index], cus.customer[index+1:]...)
		fmt.Println("----------------删除成功----------------")
	} else {
		fmt.Println("已退出.")
		return
	}
}
func (cus *CustomerList) amend(id int) {
	index := cus.findId(id)
	if index == -1 {
		fmt.Println("很遗憾,未找到该用户.")
		return
	}
	fmt.Print("你确定要修改吗？按y删除其他任意键退出.")
	tem := ""
	_, err := fmt.Scanln(&tem)
	ErrHandle(err)
	if tem == "y" || tem == "Y" {
		var name string
		fmt.Print("姓名:")
		_, err := fmt.Scanln(&name)
		ErrHandle(err)
		cus.customer[index].Name = name
		var gender string
		fmt.Print("性别:")
		_, err = fmt.Scanln(&gender)
		cus.customer[index].Gender = gender
		ErrHandle(err)
		var age int
		fmt.Print("年龄:")
		_, err = fmt.Scanln(&age)
		ErrHandle(err)
		cus.customer[index].Age = age
		var phone int
		fmt.Print("电话:")
		_, err = fmt.Scanln(&phone)
		ErrHandle(err)
		cus.customer[index].Phone = phone
		var email string
		fmt.Print("电邮:")
		_, err = fmt.Scanln(&email)
		cus.customer[index].Email = email
		fmt.Println("----------------修改成功----------------")
	} else {
		fmt.Println("已退出.")
		return
	}
}
func (cus *CustomerList) ShowDelete() {
	fmt.Println("----------------删除客户----------------")
	fmt.Print("请选择待删除客户的编号:")
	id := -1
	_, err := fmt.Scanln(&id)
	ErrHandle(err)
	cus.delete(id)
}
func (cus *CustomerList) ShowAmend() {
	fmt.Println("----------------修改客户----------------")
	fmt.Print("请选择待修改客户的编号:")
	id := -1
	_, err := fmt.Scanln(&id)
	ErrHandle(err)
	cus.amend(id)
}
func ErrHandle(err error) {
	if err != nil {
		fmt.Println("输入有误，请重新输入~~~:")
	}
}
