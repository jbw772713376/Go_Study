package main

import (
	"fmt"
	"runtime"
)

func mutatingMap_study() {
	// 定义一个map类型的变量m,并使用make初始化，cap为2
	m := make(map[string]string, 2)
	// 此时map占用的堆内存为85,修改cap的值，会导致申请的堆内存发生变化
	printMemStats()

	// 向map中添加元素
	m["a"] = "A"
	m["b"] = "B"
	fmt.Println(m, len(m))

	// 修改map中已有的元素
	m["a"] = "Aa"
	fmt.Println(m, len(m))

	// 删除map中已有的元素
	delete(m, "a")
	fmt.Println(m, len(m))

	// 判断key是否存在于map中
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("no key")
	}

}

// Alloc：     当前堆上对象占用的内存大小;
// TotalAlloc：堆上总共分配出的内存大小;
// Sys：       程序从操作系统总共申请的内存大小;
// NumGC：     垃圾回收运行的次数。
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
