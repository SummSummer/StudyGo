package main

import "unsafe"

/*
	常量：
		必须在编译器就确定其值
*/

const (
	i, f = 1, 0.123
	b = false
)

func main() {
	
	// 可以在函数代码块中定义常量，未使用的常量不会引发编译错误
	{
		const x = "abc"
	}

	// 如果显示指定类型，必须确保常量左右值类型一致
	const (
		x, y int = 99, -999
		b byte = byte(x)		// x被指定为int，需要显式转换为byte
		// n = uint(y)			-999超出uint8的范围
	)

	// 常量值也可以是编译器能计算出结果的表达式
	const (
		ptrSize = unsafe.Sizeof(uintptr(0))
		strSize = len("hello world!")
	)

	// 在常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		xx uint16 = 120
		yy
		s = "abc"
		z
	)
}