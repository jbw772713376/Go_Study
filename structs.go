package main

import "fmt"

//结构体是一个字段的结合
type isStruct struct {
	X int
	Y int
}

func main() {
	//结构体可以赋值给变量
	s := isStruct{12, 32}
	//使用.来访问结构体中的字段
	s.Y = 55
	fmt.Println(s.X, s.Y)
	fmt.Printf("%T\n", s)
	fmt.Println(isStruct{99, 88})

	//结构体可以使用指针访问
	p := &s
	//使用指针访问结构体，也不能修改字段的类型
	p.X = 1e9
	fmt.Println(s)
}
