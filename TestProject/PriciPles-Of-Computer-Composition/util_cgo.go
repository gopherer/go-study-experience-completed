package main

/*
#include<windows.h>
#include<conio.h>
#include<stdio.h>

//使用WinAPI来移动控制台光标
void gotoxy(int x,int y)
{
	COORD c;
	c.X = x;
	c.Y = y;
	SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
}

//从键盘获取一次按键，但不显示到控制台
int direct()
{
	return _getch();
}
//去掉控制台光标
void hideCursor()
{
	CONSOLE_CURSOR_INFO cci;
	cci.bVisible = FALSE;
	cci.dwSize = sizeof(cci);
	SetConsoleCursorInfo(GetStdHandle(STD_OUTPUT_HANDLE),&cci);
}
//暂停
void Pause()
{
	getchar();
}
*/
import "C" //go中可以嵌套C语言的函数

//设置控制台光标位置
func GotPosition(X int, Y int) {
	//调用C语言函数
	C.gotoxy(C.int(X), C.int(Y))
}

//无显获取键盘输入的字符
func Direction() (key int) {
	key = int(C.direct())
	return
}

//设置控制台光标隐藏
func HideCursor() {
	C.hideCursor()
}

//暂停
func Pause() {
	C.Pause()
}
