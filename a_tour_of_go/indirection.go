package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// 如果一个方法的 receiver 是一个指针类型，那么在调用方法的时候，会自动进行指针的转换，
// 会将 T 转成 &T
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // (&v).Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
