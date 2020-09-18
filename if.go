package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	//if 判断条件
	if x < 0 {
		return sqrt(-x) + "i"
	}
	//math.Sqrt(float64),用于求平方根
	//fmt.Sprint用于将
	return fmt.Sprintf("%T %g", math.Sqrt(x), math.Sqrt(x))
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	//if语句中定义的变量，作用于仅在if语句中
	//return v
	//error: undefined: v
	return lim
}

func main() {
	fmt.Printf("%T\n%s\n",sqrt(4), sqrt(-9))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 4, 10),
	)
}