package main

// for 循环的 range 形式可遍历切片或映射。
// 当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

import (
	"fmt"
	"math"
)

func initArr() [3]string {
	var arr [3]string
	arr[0] = "Hello"
	arr[1] = ","
	arr[2] = "World"
	fmt.Printf("arr[0]的内存地址是 %v\n", &arr[0])
	return arr
}

func initSlice() []int {
	sli := []int{1, 2, 4, 8, 16, 32, 64, 128}
	return sli
}

func range_study() {
	//数组的遍历
	for i, v := range initArr() {
		fmt.Println(i, v, &v)
	}
	//slice的遍历
	for i, v := range initSlice() {
		fmt.Printf("2的 %d 次方是 %d\n", i, v)
	}
	//只需要索引，可以不写value
	for i := range initSlice() {
		fmt.Printf("2的 %d 次方是 %.0f\n", i, math.Pow(2, float64(i)))
	}
	//只需要value，则需要使用"_"来占位
	for _, v := range initArr() {
		fmt.Println(v)
	}
}
