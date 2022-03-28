package main
// 切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

import "fmt"

func makeSlice_study() {
	//make用于slice，map，和channel的初始化
	//其返回由make初始化完成的数据类型的值
	a := make([]int, 5)
	makeSlicePrintSlice("a", a)

	//make会使用使用参数初始化一个slice
	b := make([]int, 0, 5)
	makeSlicePrintSlice("b", b)
	c := b[:2]
	makeSlicePrintSlice("c", c)
	d := b[2:5]
	makeSlicePrintSlice("d", d)
}

func makeSlicePrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
