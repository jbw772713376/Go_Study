package main

// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
// 接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完`v, ok := <-ch`,此时 ok 会被设置为 false。

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(500 * time.Millisecond)
	}
	// 只有发送者才能关闭信道，而接收者不能。
	// 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
	// 如果不close channel，接受者将会一直等待新的值发送
	close(c)
	// 向一个已经关闭的信道发送数据会引发程序恐慌（panic）
	// c <- 100
}

func rangeAndClosed_study() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// 循环 for i := range c 会不断从信道接收值，直到它被关闭。
	for i := range c {
		fmt.Println(i)
	}
}
