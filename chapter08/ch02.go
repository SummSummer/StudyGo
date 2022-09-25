package main

import "time"

// wait 进程退出时不会等待并发任务结束，可用通道channel阻塞，然后发出退出信号
func test01() {
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		println("gorougine done.")

		close(exit)
	}()

	println("main...")
	<-exit
	println("main exit")
}

func main() {
	
}