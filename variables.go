package main

import (
    "fmt"
)

//定义类型为float32的变量列表
var c, python float32
//当为变量赋初始值时，不需要声明变量的类型
var golang, java = true, "叶聪"

func main() {
    //定义类型为int的变量，同时为其赋值
    var i int = 92
    //使用":="的短格式声明变量，短格式声明只能出现在函数中
    //函数中声明的变量必须使用,在函数外声明的变量可以不用
    j := "04"
    //打印变量的初始值
    fmt.Println(c, golang, java, i, j)
}
