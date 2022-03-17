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
	if _, err := builder.AddFromFile("15test.glade"); err !=nil{
		fmt.Println(err)
	}
	//获取window控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy",gtk.MainQuit)
	//添加鼠标按下事件
	//BUTTON_PRESS_MASK:鼠标按下，触发信号“button-press-event"
	//BUTTON_MOTION_MASK:鼠标移动,按下任意键移动，BUTTON1（左键）BUTTON2（中间键）BUTTON3（右键）
	//默认事件未设置
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
		}else if flag == 2{
			fmt.Println("按下为中间键")
		}else if flag == 3{
			fmt.Println("按下为右键")
		}

		if event.Type == int(gdk.BUTTON_PRESS){
			fmt.Println("单击")
		}else if event.Type == int(gdk.BUTTON2_PRESS){
			fmt.Println("双击")
		}
		fmt.Println("相对与窗口",event.X,event.Y)
		fmt.Println("相对与屏幕",event.XRoot,event.YRoot)
	})
	//"motion-notify-event" 按住鼠标移动触发
	win.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		//获取鼠标按下结构体变脸，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
		fmt.Println("=======相对于窗口",event.X,event.Y)
		fmt.Println("=======相对于屏幕",event.XRoot,event.YRoot)
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
