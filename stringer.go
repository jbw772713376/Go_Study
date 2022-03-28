package main
// Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type A interface {
	String() string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringer_study() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

// 当对于一个interface{}进行switch时，将会判断该interface是否有判断列表中的method实现
func interfaceSwitch_study() {
	var a interface{}
	a = Person{"Arthur Dent", 42}
	switch v := a.(type) {
	case A:
		fmt.Println(11)
		fmt.Printf("%T", v)
	default:
		fmt.Println(12)
	}
}
