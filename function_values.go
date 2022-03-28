package main
// 函数也是值。它们可以像其它值一样传递。
// 函数值可以用作函数的参数或返回值。

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionValues_study() {
	//定义一个名为 hypot 的函数
	hypot := func(x, y float64) float64 {
		return x + y
	}
	// 直接执行函数 hypot
	fmt.Println(hypot(7, 8))

	// 将函数值 hypot 作为参数传入其他的函数
	fmt.Println(compute(hypot))
	// 将函数值作为返回传入其他函数
	fmt.Println(compute(math.Pow))
}
