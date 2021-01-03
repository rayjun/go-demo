package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))

	count := count(c1)

	fmt.Printf("%d", count)
}

func count(nums [32]uint8) int {
	count := 0
	for _, x := range nums {
		for x != 0 {
			x &= x - 1
			count++
		}
	}
	return count
}
