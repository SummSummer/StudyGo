package main

// switch语句中无需显式执行break，如果想顺序执行后续的case
// switch不会优先执行default，只有所有case匹配失败后才会执行
// 需要执行fallthrough，但这样不会匹配后续条件表达式
func main() {
	switch x := 5; x {
		default:
			println(x)
		case 5:
			x += 10
			println(x)

			fallthrough
		case 6:
			x += 20
			println(x)
	}
}

// fallthrough必须在case块末尾，可以使用break阻止
func test01() {
	switch x := 5; x {
		case 5:
			x += 10
			println(x)

			if x >= 15 {
				break
			}

			fallthrough
		case 6:
			x += 20
			println(x)
	}
}