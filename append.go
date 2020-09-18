package main

import "fmt"

func main() {
	var s []int
	printSlice("s", s)

	//append 的结果是一个包含原slice所有元素加上新添加的元素的slice
	s = append(s, 0)
	printSlice("s", s)
	p := &s[0]
	fmt.Println(p)

	s = append(s, 1)
	printSlice("s", s)
	p = &s[0]
	fmt.Println(p)

	//当开始slice分配的底层数组过小，则会重新分配一个新的数组
	s = append(s, 2, 3)
	printSlice("s", s)
	p = &s[0]
	fmt.Println(p)

	s = append(s, 4)
	printSlice("s", s)
	p = &s[0]
	fmt.Println(p)
}

func printSlice(a string, x []int) {
	fmt.Printf("len is %d, cap is %d, slice is %v, %v\n", len(x), cap(x), a, x)
}