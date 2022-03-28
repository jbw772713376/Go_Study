package main
// 类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。

import (
	"fmt"
)

func typeSwitch_study() {
	doType(1996)
	doType("四月")
	doType(1996.4)
}

// 类型选择 是一种按顺序从几个类型断言中选择分支的结构。
// 类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。
// 此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。
func doType(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Printf("%v,%T\n", v, v)
	case string:
		fmt.Printf("%v,%T\n", v, v)
	default:
		fmt.Printf("%T", v)
	}
}
