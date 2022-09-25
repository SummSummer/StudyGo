package main

import (
	"fmt"
	"unsafe"
)

// 空结构 struct{}
// 无论作为数组元素类型还是本身，其长度都为零
func test01() {
	var a struct{}
	var b [100]struct{}

	println(unsafe.Sizeof(a), unsafe.Sizeof(b))	// 0 0
}

// 尽管没有分配数组内存，但依旧可以操作元素，对于切片len cap属性正常
func test02() {
	var b [100]struct{}
	s := b[:]

	b[1] = struct{}{}
	b[2] = struct{}{}

	fmt.Println(s[3], len(s), cap(s))	// {} 100 100
}

// 空结构可作为通道元素类型，用于事件通知
func test03() {
	exit := make(chan struct{})

	go func() {
		println("hello world")
		exit <- struct{}{}
	}()

	<- exit
	println("end .")
}

func main() {
	test01()
}