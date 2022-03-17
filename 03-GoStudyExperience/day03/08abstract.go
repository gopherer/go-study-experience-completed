package main

import "fmt"

type Account struct {
	AccountNum string
	Pwd string
	Balance float64
}

func (account *Account)Deposit(money float64,pwd string){
	if pwd != account.Pwd{
		fmt.Println("输入密码有误，请重新输入:")
		return
	}
	if money <= 0{
		fmt.Println("存入金额有误，请重新输入：")
	}
	account.Balance += money
}
func (account *Account)WithDraw(money float64,pwd string) {
	if pwd != account.Pwd{
		fmt.Println("输入密码有误，请重新输入：")
		return
	}
	if money <=0 || money > account.Balance{
		fmt.Println("取走金额有误，请重新输入：")
	}
	account.Balance -= money
}
func (account *Account)Query(pwd string){
	if pwd != account.Pwd{
		fmt.Println("输入密码有误，请重新输入：")
		return
	}
	fmt.Printf("账号为%v 余额为%v\n",account.AccountNum,account.Balance)
}
func main08() {
	var account  = Account{
		"gs111",
		"666666",
		100,
	}
	account.Query("666666")
	account.Deposit(100.0,"666666")
	account.Query("666666")

	account.WithDraw(50,"666666")
	account.Query("666666")


}
