package main

import "fmt"

const (
	//定义常量时没有声明类型，那么常量的类型由下文决定
	Big = 1 << 10
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

func main() {
	fmt.Println(Big, Small)
	fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}