package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isSameChar("ray", "yar"))
}

func isSameChar(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		if !strings.Contains(s2, string(v)) {
			return false
		}
	}

	return true
}
