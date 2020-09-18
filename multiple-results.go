package main

import (
    "fmt"
)

//与参数的规则一样，在声明返回值时也同样可以将多个相同类型的变量合并
//q、w、e与x、y、z为6个在相同作用域的变量
func swap_1(q, w, e string) (x, y, z string) {
    //返回参数中的变量 
    return q, w, e
}

func swap_2(Q, W, E string) (X, Y, Z string) {
    //为X,Y,Z变量赋值
    X = "hello"
    Y = "world"
    Z = "~"
    //return返回值中声明的变量
    //返回值中有变量的定义则不用再次声明需要返回的变量名
    return
}

func main() {
    //调用swap_1对a、b、c赋值
    a, b, c := swap_1("Hello", "World", "!")
    //调用swap_2对A、B、C赋值
    A, B , C := swap_2("Hello", "World", "!")
    fmt.Println(a, b, c)
    fmt.Println(A, B, C)
}