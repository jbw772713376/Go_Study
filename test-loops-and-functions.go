package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 9.0
	for i := 1; i <= 10; i++ {
		z = z - (math.Pow(z, 2)-x)/2*z
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(49))
}