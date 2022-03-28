package main

import(
    "fmt"
    "time"
)

func closure_study() {
    s := []string{"a", "b", "c"}
    for _, v := range s {
        go func(v string) {
            fmt.Println(v)
        }(v)
    }
    time.Sleep(time.Second * 1)
}
