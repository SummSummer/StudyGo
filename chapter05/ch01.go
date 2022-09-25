package main

import (
	"fmt"
	"unsafe"
)

// 字符串
type stringStruct struct {
	str unsafe.Pointer
	len int
}

func main() {
	test01()
}

// 使用`定义不做转义处理的原始字符串
func test01() {
	s := `line\r\n,
	line2`

	println(s)
}

// 遍历字符串分为byte和rune两种方式
func test02() {
	s := "asd"

	// byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: [%c]\n", i, s[i])
	}

	// rune 返回字符索引及Unicode字符
	for i, c := range s {
		fmt.Printf("%d: [%c]\n", i, c)
	}
}

// 修改字符串，需要将字符串转为可变类型([]byte或[]rune)，都会重新分配内存，并复制数据
func test03() {
	s := "Hello World!"

	bs := []byte(s)
	s2 := string(bs)

	rs := []rune(s)
	s3 := string(rs)

	println(s2, s3)
}

// 使用append函数，可将string直接追加到[]byte内
func test04() {
	var bs []byte

	bs = append(bs, "abc"...)

	fmt.Println(bs)		// [97, 98, 99]
}