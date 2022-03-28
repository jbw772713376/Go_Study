package main

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
// 在此例中，Sum 方法拥有一个名为 v，类型为 Vertex 的接收者。

import "fmt"

type Vertex struct {
	X, Y int64
}

func (v Vertex) Sum() int64 {
	return v.X + v.Y
}

// 方法只是个带接收者参数的函数。实际执行并没有不同
func funcSum(v Vertex) int64 {
	return v.X + v.Y
}

func method_study() {
	v := Vertex{2, 3}
	fmt.Println(v.Sum())
}

// func method_study() {
// 	v := Vertex{2, 3}
// 	fmt.Println(funcSum(v))
// }
