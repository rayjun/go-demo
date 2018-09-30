package main


import (
	"golang.org/x/tour/wc"
	"strings"
)


func WordCount(s string) map[string]int {

	sat := make(map[string]int)
	strs := strings.Split(s, " ")

	for _,k := range strs {

			_, ok := sat[k]

			if ok == true {
				sat[k] += 1
			}else {
				sat[k] = 1
			}
    }
	return sat
}


func main() {
	wc.Test(WordCount)
}