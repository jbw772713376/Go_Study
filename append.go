package main
// 为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。

import "fmt"

func append_study() {
	var s []int
	appendPrintSlice("s", s)

	//append 的结果是一个包含原slice所有元素加上新添加的元素的slice
	s = append(s, 0)
	appendPrintSlice("s", s)
	p := &s[0]
	fmt.Println(p)

	s = append(s, 1)
	appendPrintSlice("s", s)
	p = &s[0]
	fmt.Println(p)

	//当开始slice分配的底层数组过小，则会重新分配一个新的数组
	s = append(s, 2, 3)
	appendPrintSlice("s", s)
	p = &s[0]
	fmt.Println(p)

	s = append(s, 4)
	appendPrintSlice("s", s)
	p = &s[0]
	fmt.Println(p)
}

func appendPrintSlice(a string, x []int) {
	fmt.Printf("len is %d, cap is %d, slice is %v, %v\n", len(x), cap(x), a, x)
}