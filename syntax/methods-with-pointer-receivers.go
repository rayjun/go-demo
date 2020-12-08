package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 所有的方法都应该使用方法指针，使用指针 receiver 的原因有两点：
// 1. 方法中可以直接修改 receiver 中的值
// 2. 可以避免 struct 的值 在每个方法中进行拷贝，特别是当 struct 很大的时候
// 最好不要两种方式混用
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
