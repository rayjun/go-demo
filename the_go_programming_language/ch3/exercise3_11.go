package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var start, end int
	if strings.Contains(s, ".") {
		end = strings.Index(s, ".")
	} else {
		end = len(s)
	}

	var buffer bytes.Buffer
	if strings.HasPrefix(s, "-") {
		start = end % 3
		if start == 1 {
			start += 3
		}
	} else {
		start = end % 3
	}

	if end <= 3 {
		return s
	}

	buffer.WriteString(s[:start])

	for i := start; i < end; i += 3 {
		buffer.WriteString("," + s[i:i+3])
	}

	buffer.WriteString(s[end:])
	return buffer.String()
}

func main() {
	fmt.Println(comma("-123.567890"))
}
