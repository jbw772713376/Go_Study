package main
// Go 只有一种循环结构：for 循环。
// 基本的 for 循环由三部分组成，它们用分号隔开：
	// 初始化语句：在第一次迭代前执行(可选)
	// 条件表达式：在每次迭代前求值
	// 后置语句：在每次迭代的结尾执行(可选)
// 初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。
// 一旦条件表达式的布尔值为 false，循环迭代就会终止。

import "fmt"

func for_study() {
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

	// 去掉 ; 也可以，此时这个就是 while
	ADD := 1
	for ADD < 10 {
		ADD += ADD
		fmt.Println(ADD)
	}
}