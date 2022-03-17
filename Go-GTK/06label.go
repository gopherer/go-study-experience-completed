package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _,err := builder.AddFromFile("06test.glade"); err != nil{
		fmt.Println(err)
	}
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//获取标签控件
	label := gtk.LabelFromObject(builder.GetObject("label1"))
	label.SetText("hello label")//设置内容
	label.ModifyFontSize(10)//设置字体大小
	label.ModifyFG(gtk.STATE_NORMAL,gdk.NewColor("red"))//设置字体颜色
	//获取标签内容
	str := label.GetText()
	fmt.Println("str = ",str)
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
