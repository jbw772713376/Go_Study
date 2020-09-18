package main

import (
    "fmt"
    "math/cmplx"
)

var (
    //ToBe bool类型为true或者false
    ToBe bool = true
    // MaxInt uint64类型
    MaxInt uint64 = 1<<64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
}

//简单类型：int、float、complex、bool和string
    //int8、int16、int32、int64，8、16、32、64表示占用的bit数，所以int最大可以存放的值为2^64-1位的整数
    //uint,无符号的整数
    //float是浮点数，建议使用float64，计算结果更加准确

//数据结构或组合类型：struct、array、slice、map和channel
//接口：interface
