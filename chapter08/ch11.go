package main

import "sync"

// 单向通道
func test01() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	c := make(chan int)
	var send chan <- int = c
	var recv <- chan int = c

	go func() {
		defer wg.Done()

		for x := range recv {
			println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	wg.Wait()

	/*
		不能在单向通道上做逆向操作
		<- send
		recv <- 1
		close(recv)	接收端不能close
		不能转换send和recv
	*/
}

func main() {
	
}