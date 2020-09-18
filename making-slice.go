package main

import "fmt"

func main() {
	//make用于slice，map，和channel的初始化
	//其返回由make初始化完成的数据类型的值
	a := make([]int, 5)
	printSlice("a", a)

	//make会使用使用参数初始化一个slice
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := b[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
