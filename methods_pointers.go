package main

// 指针可以作为接收者声明方法。

import (
	"fmt"
	"math"
)

type Methods_pointers struct {
	X, Y float64
}

func (v Methods_pointers) methodsPointersAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 若不使用指针,则方法将对 v 的副本进行操作。只有使用指针，才会对变量进行修改。
// 指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
func (v *Methods_pointers) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Methods_pointers) Scaletest(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// v 并非指针，但是带指针接收者的方法也能被直接调用,方法对于指针接受者,将会解释为 (&v).Scale(5)
func methods_pointers_study() {
	v := Methods_pointers{3, 4}
	v.Scale(10)
	// v.Scaletest(10)
	// (&v).Scale(10)
	fmt.Println(v.methodsPointersAbs())
}
