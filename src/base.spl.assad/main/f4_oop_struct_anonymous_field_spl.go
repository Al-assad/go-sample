package main

import (
	"fmt"
)

// @desc go struct 匿名字段、匿名内嵌结构体 示例

// 外部结构体
type outerS struct {
	b      int
	c      float32
	int    // 匿名字段
	innerS // 使用结构体作为内嵌字段
}

//  内嵌结构体
type innerS struct {
	in1 int
	in2 int
}

func main() {
	outer := new(outerS)
	outer.b = 10
	outer.c = 20.3
	outer.int = 30 // 匿名字段赋值
	outer.in1 = 40 // 匿名内嵌结构体赋值
	outer.in2 = 50

	fmt.Printf("%v\n", *outer) // {10 20.3 30 {40 50}}
	// 访问匿名字段
	fmt.Println(outer.int) // 30
	// 访问内嵌结构体
	fmt.Println(outer.in1, outer.in2) // 40 50
}
