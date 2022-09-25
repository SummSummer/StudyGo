package main

import (
	"fmt"
	"unsafe"
)

// 通道

// 可用于事件通知
func test01() {
	done := make(chan struct{})		// 结束事件

	c := make(chan string)			// 数据传输通道

	go func() {
		s := <-c					// 接收消息
		println(s)
		close(done)					// 关闭通道
	}()

	c <- "hi!"						// 发送消息
	<-done							// 阻塞，直到数据或通道关闭
}

// 同步模式必须有配对操作的goroutine出现，否则会一直阻塞，而异步模式在缓冲区未满或数据未读完前，不会阻塞
func test02() {
	c := make(chan int, 3)			// 创建带3个缓冲槽的异步通道

	c <- 1							// 缓冲区未满，不会阻塞
	c <- 2

	println(<-c)					// 缓冲区尚有数据，不会阻塞
	println(<-c)
}

// 缓冲区大小仅是内部属性，不属于类型组成部分
func test03() {
	var a, b chan int = make(chan int, 3), make(chan int)

	var c chan bool

	println(a == b)					// false
	println(c == nil)				// true

	fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))
}

// 内置函数cap和len返回缓冲区大小和当前已缓冲数量，对于同步通道则都返回0
func test04() {
	a, b := make(chan int), make(chan int, 3)

	b <- 1
	b <- 2

	println("a:", len(a), cap(a))
	println("b:", len(b), cap(b))
}

func main() {
	test02()
}