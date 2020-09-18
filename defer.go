package main

import "fmt"

func main() {
	fmt.Println("Countinu")

	//当defer被压入一个栈中，返回顺序遵循"后进先出"原则
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
