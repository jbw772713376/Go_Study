package main

import "fmt"

//结构体是一个字段的结合
type isStruct struct {
	X int
	Y int
}

func structs_study() {
	//结构体可以赋值给变量
	s := isStruct{12, 32}
	//使用.来访问结构体中的字段
	s.Y = 55
	fmt.Println(s.X, s.Y)
	fmt.Printf("%T\n", s)
	fmt.Println(isStruct{99, 88})

	//结构体字段可以通过结构体指针来访问。
	// 如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	p := &s
	//使用指针访问结构体，也不能修改字段的类型
	p.X = 1e9
	fmt.Println(s)
}
