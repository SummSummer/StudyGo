package main

import (
	"bytes"
	"strings"
)

// 字符串的拼接

// 使用strings.Join函数，它会统计所有参数的长度，一次性完成内存分配
func test01() string {
	s := make([]string, 1000)

	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}

	return strings.Join(s, "")
}

// bytes.Buffer
func test02() string {
	var b bytes.Buffer
	b.Grow(1000)

	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}

	return b.String()
}

// rune类型是专门用来存储Unicode码点的，是int32的别名

func main() {
	
}