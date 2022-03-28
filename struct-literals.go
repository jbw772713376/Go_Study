package main
// 结构体文法通过直接列出字段的值来新分配一个结构体。
// 使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
// 特殊的前缀 & 返回一个指向结构体的指针。

import "fmt"

type structLiteralsIsStruct struct {
	X, Y int
}
//结构体可重复使用，分配给不同的变量
var (
	v1 = structLiteralsIsStruct{1,2}
	v2 = structLiteralsIsStruct{Y: 3}
	v3 = structLiteralsIsStruct{}
	v4 = &structLiteralsIsStruct{4, 5}
)

func structLiterals_study() {
	fmt.Println(v1, v2, v3, v4, *v4)
}
