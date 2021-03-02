package main

import (
	"fmt"
)

// @desc go struct 内嵌结构示例

// 外部结构体
type outerSt struct {
	a  int
	b  float64
	in innerSt // 内嵌结构体
}

//  内嵌结构体
type innerSt struct {
	in1 int
	in2 int
}

func main() {
	outer := new(outerSt)
	outer.a = 10
	outer.b = 20.3
	outer.in.in1 = 30 // 内嵌结构体字段赋值
	outer.in.in2 = 40

	fmt.Printf("%v\n", *outer) // {10 20.3 {40 50}}
	// 访问内嵌结构体
	fmt.Println(outer.in.in1, outer.in.in2) // 30 40
}
