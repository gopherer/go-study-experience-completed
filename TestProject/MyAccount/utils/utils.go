package utils

import "fmt"

type myAccount struct {
	balance float64
	money   float64
	note    string
	reason  string
}

func NewMyAccount() *myAccount {
	return &myAccount{
		balance: 1000.0,
		money:   0,
		note:    "收支\t账号金额\t收支金额\t说  明",
		reason:  "",
	}
}
func (myA *myAccount) ShowMenu() {
Flag:
	for {
		myA.menu()
		var key int
		_, err := fmt.Scanln(&key)
		myA.errHandle(err)
		switch key {
		case 1:
			myA.details()
		case 2:
			myA.income()
		case 3:
			myA.pay()
		case 4:
			myA.exit()
			break Flag
		default:
			fmt.Println("输入有误 请重新输入")
		}
	}
}
func (myA *myAccount) menu() {
	fmt.Println("----------家庭收支记账软件----------")
	fmt.Println("            1 收支明细")
	fmt.Println("            2 登记收入")
	fmt.Println("            3 登记支出")
	fmt.Println("            4 退    出")
	fmt.Print("           请选择(1-4):")
}
func (myA *myAccount) details() {
	fmt.Println("----------当前收支明细记录----------")
	if myA.reason == "" {
		fmt.Println("当前没有收支明细...来一笔吧！")
		return
	}
	fmt.Println(myA.note)
}
func (myA *myAccount) income() {
	fmt.Print("本次收入金额:")
	if _, err := fmt.Scanln(&myA.money); err != nil {
		myA.errHandle(err)
	}
	if myA.money < 0 {
		fmt.Println("存入余额有误 请重新输入")
		return
	}
	myA.balance += myA.money
	fmt.Print("本次收入说明:")
	if _, err := fmt.Scanln(&myA.reason); err != nil {
		fmt.Println("你未填入原因")
		myA.reason = "未说明原因"
		myA.note += fmt.Sprintf("\n收入\t%v    \t%v\t\t%v", myA.balance, myA.money, myA.reason)
		return
	}
	myA.note += fmt.Sprintf("\n收入\t%v    \t%v\t\t%v", myA.balance, myA.money, myA.reason)
}
func (myA *myAccount) pay() {
	fmt.Print("本次支出金额:")
	if _, err := fmt.Scanln(&myA.money); err != nil {
		myA.errHandle(err)
	}
	if myA.money > myA.balance {
		fmt.Println("支付金额大于余额 请重新输入")
		return
	}
	myA.balance -= myA.money
	fmt.Print("本次支出说明:")
	if _, err := fmt.Scanln(&myA.reason); err != nil {
		fmt.Println("你未填入原因")
		myA.reason = "未说明原因"
		myA.note += fmt.Sprintf("\n收入\t%v    \t%v\t\t%v", myA.balance, myA.money, myA.reason)
		return
	}
	myA.note += fmt.Sprintf("\n支出\t%v    \t%v\t\t%v", myA.balance, myA.money, myA.reason)
}
func (myA *myAccount) exit() {
	fmt.Println("你确定要退出吗？确定y,按任意键取消")
	var n byte
	if _, err := fmt.Scanf("%c", &n); err != nil {
		myA.errHandle(err)
	}
	if n == 'y' || n == 'Y' {
		fmt.Print("程序已退出")
		return
	}
}
func (myA *myAccount) errHandle(err error) {
	if err != nil {
		fmt.Println("输入有误，请重新输入")
		return
	}
}
