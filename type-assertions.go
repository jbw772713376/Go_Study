package main

import "fmt"
//类型断言 提供了访问接口值底层具体值的方式。
// t := i.(T)

type typyAssertions interface{}

func typyAssertions_study() {
	var i typyAssertions
	i = "hello world"

	// 该语句断言接口值 i 保存了具体类型 string，并将其底层类型为 string 的值赋予变量 s。
	s := i.(string)
	fmt.Printf("%v,%T\n", s, s)

	// 若 i 保存了一个 string，那么 s 将会是其底层值，而 ok 为 true
	s, ok := i.(string)
	fmt.Println(s, ok)

	// 否则，err 将为 false 而 f 将为 int64 类型的零值，程序并不会产生panic。
	f, err := i.(int64)
	fmt.Println(f, err)

	// 若 i 并未保存 string 类型的值，该语句就会触发一个 panic。
	f = i.(int64) // panic
	fmt.Println(f)
}
