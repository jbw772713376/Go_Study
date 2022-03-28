package main
// switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。

import (
	"fmt"
	"runtime"
)

func switch_study() {
	fmt.Printf("%s", "Go run on: ")
	//switch 变量，匹配到则不再向下匹配
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}