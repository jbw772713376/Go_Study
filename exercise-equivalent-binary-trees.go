package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	//ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		ch1 <- 0
	}()
	go func() {
		Walk(t2, ch2)
		ch2 <- 0
	}()

	for {
		v1, v2 := <-ch1, <-ch2
		fmt.Println(v1, v2)
		if v1 == 0 && v2 == 0 {
			return true
		} else if v1 == v2 {
			continue
		} else {
			return false
		}
	}
}

func binaryTree_test() {
	t1 := tree.New(3)
	t2 := tree.New(3)
	fmt.Println(Same(t1, t2))
}
