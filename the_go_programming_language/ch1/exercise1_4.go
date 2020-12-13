package main

import (
	"bufio"
	"fmt"
	"os"
)

// 对于那些文件中出现重复行的文件，将其名称记录并打印出来
func main() {
	counts := make(map[string]int)
	fileCounts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, fileCounts)
			f.Close()
		}
	}
	for line, n := range fileCounts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCounts map[string]int) {
	input := bufio.NewScanner(f)
	isRe := false
	for input.Scan() {
		if counts[input.Text()] != 0 {
			isRe = true
		}
		counts[input.Text()]++
	}

	// 记录有重复行的文件
	if isRe {
		fileCounts[f.Name()]++
	}
}
