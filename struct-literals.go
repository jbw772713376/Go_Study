package main

import "fmt"

type isStruct struct {
	X, Y int
}
//结构体可重复使用，分配给不同的变量
var (
	v1 = isStruct{1,2}
	v2 = isStruct{Y: 3}
	v3 = isStruct{}
	v4 = &isStruct{4, 5}
)

func main() {
	fmt.Println(v1, v2, v3, v4, *v4)
}
