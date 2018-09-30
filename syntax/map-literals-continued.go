package main

import (
	"fmt"
)


type Vertex struct {
	Lat, Long float64
}


var m = map[string]Vertex{
	"Bell Labs": {40, 74},
	"Google": {37, -122},
}


func main() {
	fmt.Println(m)
}