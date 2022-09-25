package main

import "unsafe"

// 切片
type slice struct {
	array unsafe.Pointer
	len int		// 限定可读的写元素数量
	cap int		// 切片所引用数组片段的真实长度
}

// reslice以[cap]slice作为数据源创建新的切片

// append向切片尾部[len]slice添加数据，返回新的切片对象

// copy复制两个切片对象，复制的长度以较短的为准

func main() {
	
}