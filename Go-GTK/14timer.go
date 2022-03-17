package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _ , err := builder.AddFromFile("14test.glade"); err != nil{
		fmt.Println(err)
	}
	id := 0
	num := 0
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy",gtk.MainQuit)
	buttonStart := gtk.ButtonFromObject(builder.GetObject("buttonStart"))
	buttonStop := gtk.ButtonFromObject(builder.GetObject("buttonStop"))
	buttonStop.SetSensitive(false) //按钮变灰
	label := gtk.LabelFromObject(builder.GetObject("label1"))
	label.SetText("0")
	label.ModifyFontSize(30)
	//启动定时器
	buttonStart.Clicked(func() {
		id = glib.TimeoutAdd(1000, func() bool {
			num++//时间加一
			label.SetText(strconv.Itoa(num))
			buttonStart.SetSensitive(false)
			buttonStop.SetSensitive(true)
			return true
		})
	})
	//暂定定时器
	buttonStop.Clicked(func() {
		glib.TimeoutRemove(id)
		buttonStart.SetSensitive(true)
		buttonStop.SetSensitive(false)
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
