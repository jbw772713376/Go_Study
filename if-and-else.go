package main
// 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

import (
	"fmt"
	"math"
)

func ifAndElsePow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%v -> %v\n", v, lim)
	}
	return lim
}

func ifAndElse_study() {
	fmt.Println(
		ifAndElsePow(3, 2, 10),
		ifAndElsePow(3, 3, 10),
	)
}
