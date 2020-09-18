package main

import "fmt"

func main() {
	sum := 0
	// for 初始化循环变量; 循环条件； 循环变量表达式
	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	SUM := 1
	// for ; 循环条件;
	for ; SUM < 10; {
		SUM += SUM
		fmt.Println(SUM)
	}
}