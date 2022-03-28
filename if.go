package main
// Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
// 该语句声明的变量作用域仅在 if 之内。

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

func ifPow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	//if语句中定义的变量，作用于仅在if语句中
	//return v
	//error: undefined: v
	return lim
}

func if_study() {
	fmt.Printf("%T\n%s\n",sqrt(4), sqrt(-9))
	fmt.Println(
		ifPow(3, 2, 10),
		ifPow(3, 4, 10),
	)
}