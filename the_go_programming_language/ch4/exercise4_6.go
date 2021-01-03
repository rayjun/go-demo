package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	bts := []byte("ray  jun hello")
	bts = deleteSpace(bts)
	fmt.Println(string(bts))
}

// 为什么不能在 append 不能用在 []byte 上？
func deleteSpace(b []byte) []byte {
	var buf bytes.Buffer

	for i := 0; i < len(b); {
		r, size := utf8.DecodeRuneInString(string(b[i:]))
		if unicode.IsSpace(r) {
			nextrune, _ := utf8.DecodeRuneInString(string(b[i+size:]))
			if !unicode.IsSpace(nextrune) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}

		i += size
	}

	return buf.Bytes()
}
