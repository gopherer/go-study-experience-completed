package view

import (
	"Project/customerManage/service"
	"fmt"
)

type CustomerView struct {
	key          int
	loop         bool
	customerList service.CustomerList
}

func (cus CustomerView) MainView() {
	for !cus.loop {
		cus.view()
		_, err := fmt.Scanln(&cus.key)
		service.ErrHandle(err)
		switch cus.key {
		case 1:
			cus.customerList.Add()
		case 2:
			cus.customerList.ShowAmend()
		case 3:
			cus.customerList.ShowDelete()
		case 4:
			cus.customerList.ShowList()
		case 5:
			cus.exit()
		default:
			fmt.Println("输入有误，请重新输入:")
		}
	}
}
func (cus *CustomerView) view() {
	fmt.Println("----------客户信息管理软件----------")
	fmt.Println("	  1添 加 客 户")
	fmt.Println("	  2修 改 客 户")
	fmt.Println("	  3删 除 客 户")
	fmt.Println("	  4客 户 列 表")
	fmt.Println("	  5退       出")
	fmt.Print("        请选择(1-5):")
}
func (cus *CustomerView) exit() {
	fmt.Print("你是否要退出？按y退出其余任意键取消:")
	var t string
	_, err := fmt.Scanln(&t)
	service.ErrHandle(err)
	if t == "Y" || t == "y" {
		cus.loop = true
		fmt.Print("        程序已退出")
	}
}
