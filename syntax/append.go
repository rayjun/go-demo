package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlices(s)

	s = append(s, 0)
	printSlices(s)

	s = append(s,1)
	printSlices(s)


	s = append(s, 2,3,4,5,6)
	printSlices(s)
}


func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s),s)
}