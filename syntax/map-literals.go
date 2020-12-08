package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map 字面量，key 是必须的
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -7439967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
