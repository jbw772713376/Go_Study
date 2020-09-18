package main

import "fmt"

func main() {
	i, j := 11, 22

	//&符号生成指向变量内存的指针
	x := &i
	//*获取指针指向的内存，空值为nil
	fmt.Println(*x)
	*x = 33
	fmt.Println(i)

	x = &j
	*x = *x / 2
	fmt.Println(j)
}
