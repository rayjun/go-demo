package main

import (
	"fmt"
)

const Pi = 3.14

// 常量的声明和普通的变量一样，但是要使用 const 关键字
// 常量的类型只能是字符，字符串，布尔值，数值
// 常量赋值不能使用 :=
func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
