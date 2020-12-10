package main

import "fmt"

// struct 中的属性也可以通过指针来访问，正常情况下需要通过 (*p).X 来访问，但是这样台累赘， Go 允许直接使用 p.X 来访问(这样不会出现混乱吗？)
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
