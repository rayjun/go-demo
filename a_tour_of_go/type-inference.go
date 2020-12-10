package main

import "fmt"

// 当使用 := 进行赋值时，即使没有明确指明右边的类型，会根据右边的值，匹配一个最合适的类型， 比如:
// i := 42 // int
// f := 3.142 // float64
// g := 0.867 + 0.5i //complex128
func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}
