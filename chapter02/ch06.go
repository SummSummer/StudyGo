package main

// 别名，别名类型无需转换，可直接赋值

/*
	官方的两个别名
	byte	alias for uint8
	rune	alias for int32
*/

func test(x byte) {
	println(x)
}

func main() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	test(c)
}

// 并不是底层结构相同就属于别名，在64位平台上，int和int64结构完全相同，但属于不同类型，需要显式转换
func test02() {
	var x int = 4
	var y int64 = int64(x)
	println(x, y)
}