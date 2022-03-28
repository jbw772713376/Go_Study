package main

// select 语句使一个 Go 程可以等待多个通信操作。
// 需要注意的是，nil channel上的操作会一直被阻塞，如果没有default case,只有nil channel的select会一直被阻塞。

import "fmt"

func selectFibonacci(c, quit chan int) {
	x, y := 0, 1
	// select语句和switch语句一样，它不是循环，它只会选择一个case来处理，如果想一直处理channel，你可以在外面加一个无限的for循环
	for {
		// 如果有同时多个case去处理,比如同时有多个channel可以接收数据，那么Go会伪随机的选择一个case处理(pseudo-random)。
		select {
		case c <- x:
			x, y = y, x+y
		// <-quit 表示从ch中读取一个值,这里表示quit这个channel中有值
		case <-quit:
			fmt.Println("quit")
			// select会被return、break关键字中断：return是退出整个函数，break是退出当前select。
			return
		}
	}
}

func select_stduy() {
	c := make(chan int)
	quit := make(chan int)
	// 匿名函数是没有名称的函数。一般匿名函数嵌套在函数内部，或者赋值给一个变量，或者作为一个表达式。
	// 如果给匿名函数的定义语句后面加上()，表示声明这个匿名函数的同时并执行
	// 其中func(c string)表示匿名函数的参数，func(m string){}(msg)的msg表示传递msg变量给匿名函数，并执行。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	selectFibonacci(c, quit)
}
