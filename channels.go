package main

// 信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
// （“箭头”就是数据流的方向。）

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func channels_stduy() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
