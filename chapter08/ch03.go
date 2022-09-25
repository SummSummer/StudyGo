package main

import (
	"sync"
	"time"
)

// 如果等待多个任务结束，推荐使用sync.WaitGroup。通过设定计数器，让每个goroutine在退出前递减
func test01() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second)
			println("goroutine", id, "done.")
		}(i)
	}

	println("main...")
	wg.Wait()
	println("main exit.")
}

func main() {
	
}