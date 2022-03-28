package main

// 当 select 中的其它分支都没有准备好时，default 分支就会执行。

import (
	"fmt"
	"time"
)

func defaultSelect_study() {
	// time.Tick 将返回一个channel，可以使用range来生成一个定时器，这比Sleep()更加优雅
	// 将返回一个Time.time类型的数据
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		// defalut语句是可选的，不允许fall through行为，但允许case语句块为空块。
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
