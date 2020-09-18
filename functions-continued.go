package main

import (
    //可以对一个导入的包定义别名，在后续的使用中原包名无法使用，
    //这相当于导入了一个与原包不一样的包，所以原包还可以再次导入，否则相同的包无法导入两次
    "fmt"
    my_fmt "fmt"
)

//当两个参数类型相同时，除最后的一个参数需要声明类型，其他的参数使用逗号隔开即可
func add(x, y int) int {
    return x + y
}

func main() {
    //使用fmt的别名my_fmt包
    my_fmt.Println(add(66, 33))
    //使用fmt包
    fmt.Println("This is fmt output")
}
