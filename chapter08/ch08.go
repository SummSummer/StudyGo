package main

import (
	"fmt"
	"runtime"
	"time"
)

// Goexit 立即终止当前任务，运行时确保所以已注册延迟调用被执行
// 该函数不会影响其他并发任务，不会引发panic，自然也就无法捕获
func test01() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)							// 执行
		defer println("a")							// 执行

		func() {
			defer func() {
				println("b", recover() == nil)		// 执行 recover返回nil
			}()

			func() {
				println("c")
				runtime.Goexit()					// 立即终止整个调用堆栈，之后全部不执行
				println("c done.")
			}()

			println("b done.")
		}()

		println("a done.")
	}()

	<-exit
	println(exit)
}

// 如果在main里调用Goexit，它会等待其他任务结束，然后让进程直接崩溃
func main() {
	for i := 0; i < 2; i++ {
		go func(x int) {
			for n := 0; n < 2; n++ {
				fmt.Printf("%c: %d\n", 'a'+x, n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	runtime.Goexit()
	println("main exit.")
}