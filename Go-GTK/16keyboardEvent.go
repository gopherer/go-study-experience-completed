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
	if _, err := builder.AddFromFile("16test.glade"); err !=nil{
		fmt.Println(err)
	}
	//获取window控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy",gtk.MainQuit)
	//"key-press-event" 键盘按下是触发
	win.Connect("key-press-event",func(ctx *glib.CallbackContext){
		//获取键盘按下的属性结构体变量，系统内部的变脸，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventKey)(unsafe.Pointer(&arg))
		//获取到实际三字母ascii
		key := event.Keyval
		if key == gdk.KEY_a{
			fmt.Println("a")
		}else if key == gdk.KEY_b{
			fmt.Println("b")
		}else if key == gdk.KEY_c{
			fmt.Println("c")
		}
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
