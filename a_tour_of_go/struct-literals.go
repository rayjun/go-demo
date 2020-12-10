package main

import "fmt"

// struct 字面量意味着为这些属性新分配了内存
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // v1 是 Vertex 类型
	v2 = Vertex{X: 1}  // Vertex 类型，Y 为 0
	v3 = Vertex{}      // Vertex 类型，X 为 0， Y 为 0
	p  = &Vertex{1, 2} // *Vertex 类型
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
