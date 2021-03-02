package main

import "fmt"

// 函数回调示例，使用其他函数作为入参

func main() {
	compute(233, 104, add2)
	compute(233, 104, subtract2)
	compute(233, 104, multiply2)
	compute(233, 104, divide2)
}

// 回调函数，how 使用函数作为入参
func compute(a int, b int, how func(int, int)) {
	how(a, b)
}

func add2(a, b int) {
	fmt.Println("+", a+b)
}

func subtract2(a, b int) {
	fmt.Println("-", a-b)
}

func multiply2(a, b int) {
	fmt.Println("*", a*b)
}

func divide2(a, b int) {
	fmt.Println("/", a/b)
}
