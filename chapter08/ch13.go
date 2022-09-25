package main

import (
	"fmt"
	"time"
)
// 当所以通道都不可用时，select会执行default语句，如此可避开select阻塞，但需注意处理外层循环，以免陷入空耗
func test01() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			select {
			case x, ok := <- c:
				if !ok {
					return
				}
				fmt.Println("data:", x)
			default:
			}

			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	c <- 100
	close(c)

	<- done
}

// 也可以用default处理一些默认逻辑
func test02() {
	done := make(chan struct{})

	data := []chan int {
		make(chan int, 3),
	}

	go func() {
		defer close(done)

		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i:
			default:
				data = append(data, make(chan int, 3))
			}
		}
	}()

	<- done

	for i := 0; i < len(data); i++ {
		c := data[i]
		close(c)

		for x := range c {
			print(x)
		}
		println()
	}
}

func main() {
	test02()
}