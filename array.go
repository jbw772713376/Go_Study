package main
// 类型 [n]T 表示拥有 n 个 T 类型的值的数组。
// 数组的长度是其类型的一部分，因此数组不能改变大小。

import "fmt"

func array_study() {
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
