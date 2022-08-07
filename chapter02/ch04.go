package main

/*
	枚举
		可借助iota标识符实现一组自增常量值来实现枚举
*/
func main() {
	const (
		x = iota
		y
		z
	)

	const (
		_ = iota
		KB = 1 << (10 * iota)		// 1 << (10 * 1)
		MB							// 1 << (10 * 2)
		GB							// 1 << (10 * 3)
	)

	// 每行可以使用多个iota，它们会单独计数，但需要保证每行的常量个数相同
	const (
		_, _ = iota, iota * 10		// 0, 0 * 10
		a, b						// 1, 1 * 10
		c, d						// 2, 2 * 10
	)
}

func test() {
	// 如果中断了iota，必须显式恢复，且后续自增值按行序递增
	const (
		a = iota					// 0
		b							// 1
		c = 100						// 100
		d							// 100
		e = iota					// 4 (恢复iota，但计数包括c, d)
		f							// 5
	)
}

func test01() {
	// 自增默认为int，可以显示指定类型
	const (
		a = iota					// int
		b float32 = iota			// float32
		c = iota					// int (如果不显式指定iota，则与b的数据类型相同)
	)
}