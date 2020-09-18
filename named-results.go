package main

import(
    "fmt"
)

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - 1 + 3
    //当返回值已经被命名时，return时无需再次声明
    return
}

func main() {
    fmt.Println(split(12))
}
