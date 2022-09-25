package main

import (
	"fmt"
	"reflect"
)

// 字段标签 tag
// 用于对字段进行描述的元数据，在运行期可通过反射获取标签信息，常用于数据校验，数据库关系映射等
type user struct {
	name string `昵称`
	sex byte `性别`
} 

func test01() {
	u := user{"Tom", 1}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}
}

func main() {
	test01()
}