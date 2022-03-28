package main

// 映射将键映射到值。
// 映射的零值为 nil 。nil 映射既没有键，也不能添加键。
// make 函数会返回给定类型的映射，并将其初始化备用。

import "fmt"

type vertex struct {
	Lat  float64
	Long float64
}

//定义变量m的类型为map，此时m的值为nil，无法赋值
//var m map[string]vertex

//map支持在定义时，同时填充元素
var a = map[string]vertex{
	"Hello World": vertex{
		-1.23, -4.56,
	},
	"Test": vertex{
		-1.00, -2.00,
	},
}

var b = map[string]vertex{
	"Hello World": {2.33, -1.11},
	"Test":        {-2.33, 1.11},
}

func map_study() {
	// map必须由make函数来分配内存
	// make语法make(map[KeyType]ValueType, [cap])，cap并不必须，但是我们应该在初始化map的时候就为其指定一个合适的容量。
	m := make(map[string]vertex, 0)
	m["Hello World"] = vertex{1.23, 4.56}
	m["Test"] = vertex{1.00, 2.000}
	fmt.Println(m["Hello World"])
	fmt.Println(m, a, b)
	fmt.Println(len(m))
	//cap函数的参数不可以是map
	//fmt.Println(cap(m), cap(a))
}
