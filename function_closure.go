package main

// Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
// 闭包用于将对象存放在堆上，其声命周期与函数生命周期相同。
// https://www.sulinehk.com/post/golang-closure-details/

import "fmt"

func adder() func(int) int {
	sum := 0 // 作用域仅存在于该函数
	return func(x int) int {
		sum += x // sum 的作用域将会逃逸至闭包上
		return sum
	}
}

func functionClosure_study() {
	pos, neg := adder(), adder() // 两个变量都是闭包, sum 变量将会逃逸至堆中，生命周期与当前函数一致
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// 否则只能申请一个全局变量，但是这个全局变量将会存在与整个程序的生命周期中，代价过于沉重
// var clo = 0

// func adder(x int) int {
// 	sum += x
// 	return sum
// }

// func functionClosure_study() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(adder(i))
// 		fmt.Println(adder(-2*i))
// 	}
// }
