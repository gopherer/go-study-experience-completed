package main

import "Project/MyAccount/utils"

func main() {
	utils.NewMyAccount().ShowMenu()
}

/*
package main

import (
	"fmt"
)

func main() {
	var balance = 10000.0
	var money float64
	var note = "收支\t账号金额\t收支金额\t说  明"
	var reason string
Flag:
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退    出")
		fmt.Print("           请选择(1-4):")
		var key int
		count, err := fmt.Scanln(&key)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("count = %d\n", count)
			fmt.Println("请重新输入")
			continue
		}
		switch key {
		case 1:
			fmt.Println("----------当前收支明细记录----------")
			if reason == "" {
				fmt.Println("当前没有收支明细...来一笔吧！")
				continue
			}
			fmt.Println(note)
		case 2:
			fmt.Print("本次收入金额:")
			if count, err := fmt.Scanln(&money); err != nil {
				fmt.Println(err)
				fmt.Printf("count = %v\n", count)
				continue
			}
			if money < 0 {
				fmt.Println("存入余额有误 请重新输入")
				continue
			}
			balance += money
			fmt.Print("本次收入说明:")
			if _, err := fmt.Scanln(&reason); err != nil {
				fmt.Println("你未填入原因")
				reason = "未说明原因"
				note += fmt.Sprintf("\n收入\t%v    \t%v\t\t%v", balance, money, reason)
				//fmt.Printf("count = %v\n", count)
				continue
			}
			note += fmt.Sprintf("\n收入\t%v    \t%v\t\t%v", balance, money, reason)
		case 3:
			fmt.Print("本次支出金额:")
			if count, err := fmt.Scanln(&money); err != nil {
				fmt.Println(err)
				fmt.Printf("count = %v\n", count)
				continue
			}
			if money > balance {
				fmt.Println("支付金额大于余额 请重新输入")
				continue
			}
			balance -= money
			fmt.Print("本次支出说明:")
			if _, err := fmt.Scanln(&reason); err != nil {
				fmt.Println("你未填入原因")
				reason = "未说明原因"
				note += fmt.Sprintf("\n收入\t%v    \t%v\t\t%v", balance, money, reason)
				continue
			}
			note += fmt.Sprintf("\n支出\t%v    \t%v\t\t%v", balance, money, reason)
		case 4:
			fmt.Println("你确定要退出吗？确定y,按任意键取消")
			var n byte
			if _, err := fmt.Scanf("%c", &n); err != nil {
				fmt.Println(err)
				continue
			}
			if n == 'y' || n == 'Y' {
				break Flag
			}
			continue
		default:
			fmt.Println("输入有误 请重新输入")
		}
	}
	fmt.Print("程序已退出")
}
*/
