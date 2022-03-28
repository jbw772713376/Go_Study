package main

import "unsafe"

func closure1_study() {
	nextInt := intSeq("aaa")
	// println(unsafe.Pointer(&nextInt))
	println(nextInt("bbb")) // 1
    println(nextInt("ccc")) // 2
	// println(nextInt()) // 3

	// newInts := intSeq()
	// println(unsafe.Pointer(&newInts))
	// println(newInts()) // 1

}

func intSeq(s string) func(k string) int {
	i := 0
	println(s)
	println(unsafe.Pointer(&i))
	return func(k string) int {
		i += 1
		println(k)
		println(unsafe.Pointer(&i))
		return i
	}
}

