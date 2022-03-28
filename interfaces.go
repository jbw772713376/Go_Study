package main

// 接口类型 是由一组方法签名定义的集合。
// 接口类型的变量可以保存任何实现了这些方法的值。
// nil 接口值既不保存值也不保存具体类型。
// 为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个具体方法的类型。

import (
	"fmt"
	"math"
)

type InterfaceAbser interface {
	InterfaceAbs() float64
}

func interface_study() {
	 var a InterfaceAbser
	
	f := InterfaceMyFloat(-math.Sqrt2)
	v := InterfaceVertex{3, 4}

	a = f // a MyFloat 实现了 InterfaceAbser
	fmt.Println(a.InterfaceAbs())
	fmt.Printf("%T", a)

	 a = &v // a *Vertex 实现了 InterfaceAbser
	 fmt.Println(a.InterfaceAbs())

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// a = v
}

type InterfaceMyFloat float64

func (f InterfaceMyFloat) InterfaceAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type InterfaceVertex struct {
	X, Y float64
}

func (v *InterfaceVertex) InterfaceAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
