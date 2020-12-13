package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""

	begin := time.Now()

	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	end := time.Until(begin)

	fmt.Println(end)

	fmt.Println(s)

	begin = time.Now()

	s = strings.Join(os.Args[:], " ")

	end = time.Until(begin)

	fmt.Println(end)
	fmt.Println(s)
}
