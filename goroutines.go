package main

// Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 当调整两个函数的执行顺序，结果并不相同。
// 当 goroutine 在后面执行时，say() 没有输出。
func goRoutines_study() {
	go say("world")
	say("hello")
	// 当 goroutine 放在最后执行时，请确保 goroutine 能够在进程结束前完成处理。
	// go say("world")
	//time.Sleep(10000 * time.Millisecond)
}
