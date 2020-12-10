package main

import "fmt"

// slice 字面量与数组的字面量类似：
// 数组字面量： [3]bool{true, true, false}
// slice 字面量：[]bool{true, true, false}
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}
