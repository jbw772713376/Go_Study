package main
// 编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// 实现一个 rot13Reader 的方法,这个方法将会在 io.Copy 中 被 copyBuffer 函数调用。
func (self rot13Reader) Read(p []byte) (n int, err error) {
	leng, err := self.r.Read(p)
	for i := 0; i < leng; i++ {
		switch {
		case p[i] >= 'a' && p[i] < 'n':
			fallthrough
		case p[i] >= 'A' && p[i] < 'N':
			p[i] = p[i] + 13
		case p[i] >= 'n' && p[i] <= 'z':
			fallthrough
		case p[i] >= 'N' && p[i] <= 'Z':
			p[i] = p[i] - 13
		}
	}
	// 注意，err 的返回不要是用 nil 返回，这将使程序在 copyBuffer 中无限循环
	return leng, io.EOF
}

func readerTest_study() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
