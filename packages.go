package main

//这里使用import导入了三个包:"fmt"、"math/rand"、"net/http"
import (
    "fmt"
    "math/rand"
    //在"net/http"前加入"_"，可以匿名的导入一个包，匿名导入的包即使在后续的代码中没有使用也不会出发报错
    //匿名导入旨在导入某些包时使用其中的"init()"函数等初始化操作，但是在后续的代码中又无需使用该包中的其他函数
    _ "net/http"
)

func main () {
    //使用fmt包的Println函数打印rand.Intn生成的随机数。
    fmt.Println("My favorite number is", rand.Intn(200))
}
