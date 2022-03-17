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
	if _, err := builder.AddFromFile("10test.glade"); err != nil{
		fmt.Println(err)
	}
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//获取entry
	en := gtk.EntryFromObject(builder.GetObject("entry1"))
	//设置内容
	en.SetText("hello entry")
	//获取内容
	fmt.Println(en.GetText())
	//是否只读
	//en.SetEditable(false)
	//是否可见，密码模式
	en.SetVisibility(false)
	//"activate" 控件内部按回车健触发
	en.Connect("activate", func() {
		fmt.Println(en.GetText())
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
