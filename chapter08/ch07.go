package main

import "runtime"

// gosched 暂停，释放线程去执行其他任务。当前任务被放回队列，等待下次调度时恢复执行
func test01() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func() {
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)

			if i == 1 {
				runtime.Gosched()		// 让出当前线程，调度执行线程b
			}
		}
	}()

	<-exit
}

func main() {
	
}