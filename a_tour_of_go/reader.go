package main

import (
	"fmt"
	"io"
	"strings"
)

// io 包中 io.Reader 接口代表读取一个有限的数据流操作
// 标准库中对于这个接口有很多的实现，包括文件，网络连接，压缩，密码等等其他的
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err= %v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
