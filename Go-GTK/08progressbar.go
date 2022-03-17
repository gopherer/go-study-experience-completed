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
	if _, err := builder.AddFromFile("08test.glade"); err !=nil{
		fmt.Println(err)
	}
	//加载控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//获取进度条
	pgb := gtk.ProgressBarFromObject(builder.GetObject("progressbar1"))
	//设置进度0.0-1.0
	pgb.SetFraction(0.5)
	//获取进度条，并打印
	fmt.Println("vale = ",pgb.GetFraction())
	//设置文本内容
	pgb.SetText("50%")
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
