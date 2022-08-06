package main

// 可以为当前包内的任意类型定义方法
type X int

func (x *X) inc() {
	*x++
}

func main() {
	var x X
	x.inc()
	println(x)
}