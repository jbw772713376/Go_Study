package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 3, 4
	var f float64 = math.Sqrt(float64(a*a + b*b))
	//短格式类型转换
	u := uint(a)
	//长格式类型转换
	var i int = int(f)
	//%T显示变量的类型
	fmt.Printf("%T, %T, %T, %T, %T", a, b, f, u, i)
}