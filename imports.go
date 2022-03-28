package main
//此代码用圆括号组合了导入，这是“分组”形式的导入语句。

//import可以使用"()"来一次性导入多个包
//也可以分开来写，每一个import导入一个包，例如:
//import "fmt"
//import "math"
import (
    "fmt"
    "math"
)

func import_study() {
    fmt.Printf("Now you have %g problems.\n", math.Nextafter(3, 5))
}
