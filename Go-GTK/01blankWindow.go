package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//initialization
	gtk.Init(&os.Args)
	//user write code
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("go GTK")
	win.SetSizeRequest(1000,660)
	win.Show()
	//main event loop
	gtk.Main()
}
