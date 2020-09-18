package main

import "fmt"

func main() {
	//数组的声明方式
	var a [2]string

	//为数组赋值
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])

	//短格式声明并赋值
	b := [2]string{"it's", "golang"}
	fmt.Println(b)
}
