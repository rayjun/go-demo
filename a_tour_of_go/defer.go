package main

import "fmt"

// 使用 defer 修饰的代码会等到整个方法返回之后才会执行
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
