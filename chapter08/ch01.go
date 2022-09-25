package main

import "time"

// 和defer一样，goroutine会因延迟执行而立即计算并复制执行参数
var c int

func test01() int {
	c++
	return c
}

func main() {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second)
		println("go:", x, y)
	}(a, test01())

	a += 100
	println("main:", a, test01())

	time.Sleep(time.Second * 3)
}