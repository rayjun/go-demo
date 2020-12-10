package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go 语言中没有 类的概念，但是可以为 type 定义方法
// 方法就是一个 recevier 类型的函数, receiver 在 func 和方法名称之间
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
