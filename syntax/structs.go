package main

import "fmt"

// struct 是一系列属性的集合
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
