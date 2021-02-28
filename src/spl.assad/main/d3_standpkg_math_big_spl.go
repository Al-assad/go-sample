package main

import (
	"fmt"
	"math"
	"math/big"
)

// go math/big 精密计算

func main() {

	// 有理数计算
	num1 := big.NewInt(12345678900)
	num2 := big.NewInt(900)
	num3 := big.NewInt(math.MaxInt64)
	r1 := big.NewInt(1)
	// r1 = (num1 * num2 + num3) / num2
	r1.Mul(num1, num2).Add(r1, num3).Div(r1, num2)
	fmt.Printf("Big Int: %v \n", r1)

	// 物理数计算
	n1 := big.NewRat(1024, 33)
	n2 := big.NewRat(14444, 7812)
	r2 := big.NewRat(1, 1)
	r2.Mul(n1, n2).Add(r2, n1)
	fmt.Printf("Big Rat: %v \n", r2)
}
