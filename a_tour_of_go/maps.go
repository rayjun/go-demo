package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// 键值对 map，没有元素的 map 是 nil，nil map 没有key，也不能添加key，需要使用 make 创建一个 map，然后才能添加键值对
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
