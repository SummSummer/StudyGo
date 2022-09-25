package main

import (
	"fmt"
	"sync"
	"time"
)

// 在迭代期删除或新增键值是安全的
func test01() {
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}

	for k := range m {
		if k == 5 {
			m[100] = 1000
		}

		delete(m, k)
		fmt.Println(k, m)
	}
}

// 运行时会对字典并发操作进行检测，如果某个任务正在进行写操作，那么其他任务就不能执行并发操作(读，写，删)
func test02() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"] += 1
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			_ = m["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select{}
}

// 可以使用sync.RWMutex实现同步，避免读写操作同时进行
func test03() {
	var lock sync.RWMutex
	m := make(map[string]int)

	go func() {
		for {
			lock.Lock()
			m["a"] += 1
			lock.Unlock()

			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m["b"]
			lock.RUnlock()

			time.Sleep(time.Microsecond)
		}
	}()

	select{}
}

// 字典对象本身就是指针包装，传参时无须再次取地址
// 创建字典时，预先准备足够的空间，可以减少扩张时的内存分配和重新哈希操作

func main() {
	test01()
}