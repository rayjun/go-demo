package main

import "fmt"

// 被defer 修饰的代码会压入栈中，当外部方法返回的时候，再从栈中弹出执行，所以结果的顺序
// 与压入的顺序相反
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
