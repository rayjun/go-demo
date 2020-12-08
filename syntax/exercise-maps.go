package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)
	strs := strings.Fields(s)

	for _, s := range strs {
		if _, ok := countMap[s]; ok {
			countMap[s] = countMap[s] + 1
		} else {
			countMap[s] = 1
		}
	}
	return countMap
}

func main() {
	wc.Test(WordCount)
}
