package main
// defer 语句会将函数推迟到外层函数返回之后执行。
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
// 推迟的函数调用会被压入一个栈中。

import "fmt"

func defer_study() {
	fmt.Println("Countinu")

	//当defer被压入一个栈中，返回顺序遵循"后进先出"原则
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
