package main

import "fmt"

// 对于数值类型，0 是默认值
// 对于布尔类型，false 是默认值
// 对于字符串，"" 是默认值

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
