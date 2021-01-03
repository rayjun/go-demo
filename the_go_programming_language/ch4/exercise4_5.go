package main

import "fmt"

func main() {
	strs := []string{"ray", "ray", "jun", "hello", "jun"}
	strs = deleteRep(strs)
	fmt.Println(strs)
}

func deleteRep(strs []string) []string {
	cur := strs[0]
	out := strs[:1]

	for _, s := range strs {
		if s != cur {
			out = append(out, s)
			cur = s
		}
	}
	return out
}
