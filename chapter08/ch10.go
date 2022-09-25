package main

import (
	"sync"
	"time"
)

// 收发 除使用简单的发送和接收操作符外，还可用ok-idom或range模式处理数据
func test01() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			x, ok := <-c
			if !ok {
				return
			}
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)
	<- done
}

// range模式
func test02() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c {
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)
	<- done
}

// 通知可以是群体性的，也未必就是通知结束，可以是任何需要表达的事件
func test03() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			println(id, ": ready.")
			<- ready
			println(id, ":running...")
		}(i)
	}

	time.Sleep(time.Second)
	println("Ready? Go!")

	close(ready)

	wg.Wait()
}

// 对于closed或nil通道
// 向已关闭通道发送数据，引发panic
// 从已关闭通道接受数据，返回已缓冲数据或零值
// 无论收发，nil通道都会阻塞
func test04() {
	c := make(chan int, 3)

	c <- 10
	c <- 20
	close(c)

	for i := 0; i < cap(c) + 1; i++ {
		x, ok := <- c
		println(i, ":", ok, x)
	}
}

func main() {
	test03()
}