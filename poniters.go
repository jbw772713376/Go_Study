package main
// Go 拥有指针。指针保存了值的内存地址。

import "fmt"

func poniters_study() {
	i, j := 11, 22

	//&符号生成指向变量内存的指针
	x := &i // 指向 i
	//*获取指针指向的内存，空值为nil
	fmt.Println(*x, x) // *x 通过指针读取 i 的值
	*x = 33 // 通过指针设置 i 的值
	fmt.Println(i, *x,  x) // 查看 i 的值为 33

	i = 44
	fmt.Println(i, *x,  x)

	x = &j // 指向 j
	*x = *x / 2 // 通过指针对 j 进行除法运算
	fmt.Println(j, *x, x)
}
