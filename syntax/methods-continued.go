package main

import (
	"fmt"
	"math"
)

// 可以在非 struct 上声明方法，方法的定义和 recevier 要在同一个包下
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	// math.Sqrt(2) 可以简写为 math.Sqrt2
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
