package main

import (
	"io"
	"os"
	"strings"
)



type rot13Reader struct {
	r io.Reader
}



func (m rot13Reader) Read(p []byte) (int, error) {
	
	ori := make([]byte, 50)
	i, err := m.r.Read(ori)
	for index, value := range ori[:i] {
		p[index] = value
	}

	return i , err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}