package main

import (
	"fmt"
	"runtime"
)

func main() {
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