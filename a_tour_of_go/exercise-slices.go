package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	sliceY := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		sliceX := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			sliceX[j] = uint8(i * j)
		}
		sliceY[i] = sliceX
	}
	return sliceY
}

func main() {
	pic.Show(Pic)
}
