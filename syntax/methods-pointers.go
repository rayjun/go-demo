package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 可以通过 receiver 的指针来定义方法, *T, T 不能是类似原生数据类型的指针，比如 *int
// 只有通过指针才能修改到 receiver 的值，所以方法通常会在指针上定义
// 如果不通过指针定义，那么传过来的 receiver 只是原 receiver 值的 copy
// 如果把下面的 * 去掉，将无法修改原 receiver 的值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
