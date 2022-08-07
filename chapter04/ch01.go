package main


/*
	Go的函数
	无需前置声明
	不支持命名嵌套定义
	不支持重载
	不支持默认参数
	支持不定长变参
	支持多返回值
	支持命名返回值
	支持匿名函数和闭包
*/

// 函数属于第一类对象，具备相同参数和返回值列表的视为同一类型

// 第一类对象(first-class object)：指可在运行期创建，可用作函数参数或返回值，可存入变量的实体。最常见的用法是匿名函数

func hello() {
	println("hello world!")
}

func exex(f func()) {
	f()
}

func test01() {
	f := hello
	exex(f)
}

func main() {
	test01()
}