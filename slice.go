package main
// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
// 类型 []T 表示一个元素类型为 T 的切片。
// 切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

import "fmt"

func slice_study() {
	//切片可以理解为没有声明长度的数组
	//但是切片不是一个数组，而是指向一个数组的地址(类似于指针)
	s := []int{1, 2, 3, 4, 5, 6}
	s[2] = 33
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		if i != 2 {
			s[i] = s[i] * 11
		}
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	//range格式用于遍历循环slice
	for i, j := range s {
		fmt.Printf("1 + %d * 11 = %d\n", i, j)
	}

	//使用_忽略值
	for _, j := range s {
		fmt.Println(j)
	}

	for i := range s {
		fmt.Println(i)
	}

	//slice的零值为nli
	var n []int
	fmt.Println(n, len(n), cap(n))

	//可以对slice重新切片,s[lo:hi]表示从下标为lo的元素到下标为hi-1的元素，空为两端定点
	fmt.Println("s[1:4] ==", s[1:4])
	fmt.Println("s[2:] ==", s[2:])
	fmt.Println("s[:3] ==", s[:3])

	//s[lo:lo]为空，而s[lo:lo+1]则只有一个元素
	//也就是说切片后的元素有hi-lo个
	fmt.Println("s[2:2] ==", s[2:2])
	fmt.Println("s[2:3] ==", s[2:3])
}
