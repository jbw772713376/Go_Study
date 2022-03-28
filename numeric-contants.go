package main
// 数值常量是高精度的 值。

import "fmt"

const (
	//Big 定义常量时没有声明类型，那么常量的类型由下文决定
	Big = 1 << 10
	//Small >> or <<意为将二进制的位数减少或增加指定位
	Small = Big >> 99
)

func needInt(x int) int {
	//返回值为int类型
	return x * 10 + 1
}

func needFloat(y float64) float64 {
	//返回值为float64类型
	return y * 0.1 + 1
}

func numericContants_study() {
	fmt.Println(Big, Small)
	fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}