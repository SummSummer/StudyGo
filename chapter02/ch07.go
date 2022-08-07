package main

// 引用类型

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

// 使用new为引用类型分配内存时，仅分配了字典类型本身所需要的内存，并没有分配键值存储内存，无法正常使用
func newMap() map[string]int {
	p := new(map[string]int)	// new返回指针
	m := *p
	m["a"] = 1					// panic: assignment to entry in nil map
	return m
}

func main() {
	newMap()
}