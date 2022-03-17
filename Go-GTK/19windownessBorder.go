package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("19test.glade");err !=nil{
		fmt.Println(err)
	}
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy",gtk.MainQuit)
	win.SetDecorated(false)//去边框

	x , y := 0, 0

	win.SetEvents(int(gdk.BUTTON_PRESS_MASK|gdk.BUTTON1_MOTION_MASK))
	//”button-press-event"鼠标按下是触发
	win.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		//获取鼠标按下属性结构体变量，系统内部的变量，不是用户传参变量
		//写法参见官方文档
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
		//1为左键，2为中间键，3为右键
		flag := event.Button
		if flag == 1{
			fmt.Println("按下为左键")
		}else if flag == 3{
			gtk.MainQuit()
		}
		x, y = int(event.X),int(event.Y)

	})
	//"motion-notify-event" 按住鼠标移动触发
	win.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
		win.Move(int(event.XRoot)-x,int(event.YRoot)-y)
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
