package main
// 函数可以没有参数或接受多个参数。
// 在本例中，funcAdd 接受两个 int 类型的参数。
// 注意类型在变量名 之后。

import (
	//在导入包时在包前添加"."，可以在后续的代码中访问该包中的属性时无需再写报名
	. "fmt"
)

//使用func来定义一个函数，函数add中包含两个参数，int类型的'x'和'y'，
//最后定义一个int类型的返回值
//return中可以出现表达式"x + y"，但是不能出现赋值语句"z = x + y"
func funcAdd(x int, y int) int {
	return x + y
}

func func_study() {
	Println(funcAdd(66, 33))
}
