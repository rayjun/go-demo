package main

import (
	"fmt"
	"go-demo/the_go_programming_language/ch2/popcount"
	"time"
)

func main() {

	start := time.Now()
	popcount.PopCount(12617267126716267234)
	end := time.Since(start)

	fmt.Println("PopCount cost time: %s", end)

	start = time.Now()
	popcount.LoopPopCount(12617267126716267234)
	end = time.Since(start)

	fmt.Println("LoopPopCount cost time: %s", end)

	start = time.Now()
	popcount.ShiftPopCount(12617267126716267234)
	end = time.Since(start)

	fmt.Println("ShiftPopCount cost time: %s", end)

	start = time.Now()
	popcount.LSBPopCount(12617267126716267234)
	end = time.Since(start)

	fmt.Println("LSBPopCount cost time: %s", end)
}
