package main

func main14() {
	//创建一个channel,双向
	ch := make(chan int)

	//双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch //只能写，不能读
	var readCh <-chan int = ch  //只能读，不能写

	writeCh <- 666 //写
	//<- writeCh //invalid operation: <-writeCh (receive from send-only type chan<- int)
	<-readCh //读
	//readCh <- 666//invalid operation: readCh <- 666 (send to receive-only type <-chan int)

	//单向channel无法转换为双向
	//var ch1 chan int = writeCh// cannot use writeCh (type chan<- int) as type chan int in assignment
}
