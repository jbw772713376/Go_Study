package main
// 常量的声明与变量类似，只不过是使用 const 关键字。
// 常量可以是字符、字符串、布尔值或数值。

import "fmt"

//Pi 为常量,常量的赋值不可以使用":="
const Pi = 3.14

func constants_study() {
	const WORLD = "World"
	fmt.Println("Hello", WORLD, Pi)
	const Truth = true
	fmt.Println("YeCong is my lover?", Truth)
}