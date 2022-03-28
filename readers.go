package main

// Read 方法用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。

import (
	"fmt"
	"io"
	"strings"
)

func reader_study() {
	r := strings.NewReader("Hello, Reader!")
	// 这里不会陷入死循环的原因是 io.Reader 存在一个  WriteTo 方法
	// io.Copy(os.Stdout, r)

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
