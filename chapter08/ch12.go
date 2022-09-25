package main

import "sync"
// 选择 如果同时处理多个通道，可选用select语句，它会随机选择一个可用通道做收发操作
func test01() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()

		for {
			var (
				name string
				x int
				ok bool
			)

			select {
			case x, ok = <- a:
				name = "a"
			case x, ok = <- b:
				name = "b"
			}

			if !ok {
				return
			}

			println(name, x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()
	wg.Wait()
}

// 如果要等全部通道消息处理结束，可将已完成通道设置为nil，这样它就会被阻塞，不再会被select选中
func test02() {
	var wg sync.WaitGroup
	wg.Add(3)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()

		for {
			select {
			case x, ok := <- a:
				if !ok {
					a = nil
					break
				}
				println("a", x)
			case x, ok := <- b:
				if !ok {
					b = nil
					break
				}
				println("b", x)
			}

			if a == nil && b == nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i * 10
		}
	}()
	wg.Wait()
}

// 即使是同一通道，也会随机选择case执行
func test03() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		defer wg.Done()

		for {
			var v int
			var ok bool

			select {
			case v, ok = <- c:
				println("a1", v)
			case v, ok = <- c:
				println("a2", v)
			}

			if !ok {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 10; i++ {
			select {
			case c <- i:
			case c <- i * 10:
			}
		}
	}()

	wg.Wait()
}

func main() {
	test01()
}