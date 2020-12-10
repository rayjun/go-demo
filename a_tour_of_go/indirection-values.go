package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 这里的 receiver 是一个普通的类型，那么如果是被指针调用，那就会自动转成 (*T).Abs()
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // p.Abs() 等价于 (*p).Abs()
	fmt.Println(AbsFunc(*p))
}
