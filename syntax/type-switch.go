package main

import "fmt"

// type switch 可以用来判断多种类型，通过 case 来处理不同的类型
func do(i interface{}) {
	switch v := i.(type) { // 这里使用 i.(type) 关键字来获取真实的类型和值，v 与 i 实际mmianian一样
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2) // v 如果为 int 就会执行这里的代码
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v)) // v 如果为 string 就会执行这里的代码
	default:
		fmt.Printf("I don't know about type %T\n", v) // 其他类型就会执行这里的代码
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
