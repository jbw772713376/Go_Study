package main

import "fmt"

func deferTmp_study() {
	i := 100
	defer fmt.Println(definVar())
	i++
	fmt.Println(i)
	return
}

func definVar() int {
	return 200
}
