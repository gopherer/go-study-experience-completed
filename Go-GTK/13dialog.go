package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("13test.glade"); err !=nil{
		fmt.Println(err)
	}
	//获取windows控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy",gtk.MainQuit)
	//获取按钮
	b1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	b2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	//问题对话框
	b1.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win, //指定父窗口
			gtk.DIALOG_MODAL,//模态对话框
			gtk.MESSAGE_QUESTION,//问题对话框
			gtk.BUTTONS_YES_NO,//按钮
			"这是问题对话框")
		if response := dialog.Run(); response == gtk.RESPONSE_YES{
			fmt.Println("Yes")
		}else if response == gtk.RESPONSE_NO{
			fmt.Println("NO")
		}else{
			fmt.Println("dialog has closed")
		}
		dialog.Destroy()
	})
	//消息对话框
	b2.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win,//指定父窗口
			gtk.DIALOG_MODAL,//模态对话框
			gtk.MESSAGE_INFO,//消息对话框
			gtk.BUTTONS_OK,//按钮
			"这是消息对话框")//内容
		dialog.Run()
		dialog.Destroy()
	})
	//显示控件
	win.ShowAll()
	//主事件循环
	gtk.Main()

}
