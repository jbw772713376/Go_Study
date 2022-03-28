package main

// 接口也是值。它们可以像其它值一样传递。
// 接口值可以用作函数的参数或返回值。
// 在内部，接口值可以看做包含值和具体类型的元组：(value, type)
// 接口值保存了一个具体底层类型的具体值。
// 接口值调用方法时会执行其底层类型的同名方法。

import (
	"fmt"
	"math"
)

type InterfaceValueI interface {
	M()
}

type InterfaceValueT struct {
	S string
}

func (t *InterfaceValueT) M() {
	fmt.Println(t.S)
}

type InterfaceValueF float64

func (f InterfaceValueF) M() {
	fmt.Println(f)
}

func intercaseValue_study() {
	var i InterfaceValueI

	i = &InterfaceValueT{"Hello"}
	describe(i)
	i.M()

	i = InterfaceValueF(math.Pi)
	describe(i)
	i.M()
}

func describe(i InterfaceValueI) {
	fmt.Printf("(%v, %T)\n", i, i)
}
