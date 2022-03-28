package main

import (
	"fmt"
	"math"
)

type MethodWithPointerReceivers struct {
	X, Y float64
}

func (v *MethodWithPointerReceivers) methodWithPointerReceivers_Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *MethodWithPointerReceivers) methodWithPointerReceivers_Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodWithPointerReceivers_study() {
	v := &MethodWithPointerReceivers{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.methodWithPointerReceivers_Abs())
	v.methodWithPointerReceivers_Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.methodWithPointerReceivers_Abs())
}
