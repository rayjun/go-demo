package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buffer bytes.Buffer
	chNum := n % 3
	buffer.WriteString(s[:chNum])

	for i := chNum; i < n; i += 3 {
		buffer.WriteString("," + s[i:i+3])
	}

	return buffer.String()
}

func main() {
	fmt.Println(comma("1234567890"))
}
