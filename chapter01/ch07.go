package main

import "fmt"

// 将字典map作为内置类型，可直接从运行时层面获得性能优化
func main() {
	m := make(map[string]int)

	m["a"] = 1				// 添加或设置字典数据

	x, ok := m["b"]			// 使用ok-idiom获取值，可以知道key/value是否存在

	fmt.Println(x, ok)

	delete(m, "a")			// 删除元素
}