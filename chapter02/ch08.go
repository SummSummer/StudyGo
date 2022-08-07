package main

// 类型转换

// Go强制要求使用显式类型转换
func test01() {
	a := 10
	b := byte(a)
	c := a + int(b)
	println(c)
}

// Go不能将bool类型的结果当成true和false使用
// func test02() {
// 	x := 1

// 	var b bool = x		// cannot use x (variable of type int) as bool value in variable declaration
// 	println(b)
// 	if x {				// non-boolean condition in if statement

// 	}
// }

// 如果转换的目标是指针，单向通道或没有返回值的函数类型，必须使用括号，避免语法分解错误
func test03() {
	/*
		(*int)(p)
		(<-chan int)(c)
		(func())(x)

		(func() int)(x)
	*/
}

func main() {
	
}