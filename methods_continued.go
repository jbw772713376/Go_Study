package main

// 可以为非结构体类型声明方法。
// 接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) method_continued_Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func method_continued_study() {
	f := MyFloat(-math.Sqrt2) // math.Sqrt2 是一个没有类型的常量，需要使用 MyFloat 进行类型转换
	fmt.Println(f.method_continued_Abs())
}
