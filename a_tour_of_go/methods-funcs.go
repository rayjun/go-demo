package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//方法就是一个函数加一个 receiver
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
