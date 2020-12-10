package main

import "fmt"

// 空接口是指没有任何方法定义的接口，空接口可以持有很多的类型
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
