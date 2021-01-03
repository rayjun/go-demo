package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "ray jun hello"
	s = string(reverseString([]byte(s)))
	fmt.Println(s)
}

func reverseString(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseByte(b[i : i+size])
		i += size
	}
	reverseByte(b)
	return b
}

func reverseByte(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}
