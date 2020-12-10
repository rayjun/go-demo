package main

import "fmt"

// 类型断言提供了一种访问真实类型的方法
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string) // 通过这种方式不会报错，ok 返回 true 表示类型匹配，返回 false 表示类型不匹配
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 如果真实的类型与目标类型不一样，就会报错
	fmt.Println(f)
}
