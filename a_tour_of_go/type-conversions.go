package main

import (
	"fmt"
	"math"
)

// go 语言中，直接使用 T(v) 来进行类型转换。对类型的转换要求很严格，必须要明确指定类型, 类型方法 T 不能省略，否则会报错

// 转换方法1
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)

// 转换方法2
// i := 42
// f := float64(i)
// u := uint(f)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
