package main

import "fmt"

func main() {
	// 当变量在声明是没有指定类型时，变量的类型由右值的类型决定
	i := 420
	f := 3.14
	// 但是当右值中包含了数值常量时，变量的类型则取决于常量的精度
	c := 3 + 5i
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("c is of type %T\n", c)
}