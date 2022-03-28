package main
// 方法能够修改其接收者指向的值。
// 这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
// 通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。

import (
	"fmt"
)

type Methods_pointers_tmp struct {
	X, Y float64
}

func (v *Methods_pointers_tmp) indirrctionScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Methods_pointers_tmp) indirrctionScaleNotPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Methods_pointers_tmp, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFuncNotPointer(v Methods_pointers_tmp, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func indirrction_stduy() {
	// 函数的参数必须和声明是的一致，否则将会无法编译
	// 而接受者为指针的方法则不关心参数的类型，而是在编译时，将参数 v.indirrctionScale(2) 解释为 (*v).indirrctionScale(2)
	v := Methods_pointers_tmp{3, 4}
	v.indirrctionScale(2)
	ScaleFunc(&v, 10)

	p := &Methods_pointers_tmp{3, 4}
	p.indirrctionScale(2)
	ScaleFunc(p, 10)

	// 当一个方法的接受者不为指针时，接受者以值调用则会将 y.indirrctionScaleNotPointer(2) 解释为 (*y).indirrctionScaleNotPointer(2)
	// 当然，不以指针为接受者的方法，自然无法读接受者进行修改，而是针对接受者的副本进行修改
	x := Methods_pointers_tmp{3, 4}
	x.indirrctionScaleNotPointer(2)
	ScaleFuncNotPointer(x, 10)

	y := &Methods_pointers_tmp{3, 4}
	y.indirrctionScaleNotPointer(2)
	ScaleFuncNotPointer(*y, 10)

	fmt.Println(v, p, x, y)
}
