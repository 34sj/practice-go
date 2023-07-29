package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var (
		a = 10
		b = 2
		sum = 0
		sub = 0
		mul = 0
		div = 0
	)

	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
