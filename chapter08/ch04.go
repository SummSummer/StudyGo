package main

import (
	"sync"
	"time"
)

// 可在多处使用wait阻塞，它们都能接收到通知
func test01() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()					// 等待归零，解除阻塞
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done...")
		wg.Done()					// 递减计算
	}()

	wg.Wait()						// 等待归零，解除阻塞
	println("main exit...")
}

func main() {
	
}