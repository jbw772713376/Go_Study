package main

import "fmt"

//Pi 为常量,常量的赋值不可以使用":="
const Pi = 3.14

func main() {
	const WORLD = "World"
	fmt.Println("Hello", WORLD, Pi)
	const Truth = true
	fmt.Println("YeCong is my lover?", Truth)
}