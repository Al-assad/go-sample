package main

import (
	"fmt"
	"strings"
)

// 闭包（匿名函数）示例

func main() {

	// 定义闭包函数 sumFunc
	sumFunc := func() int {
		cur := 0
		for i := 0; i < 100; i++ {
			cur += i
		}
		return cur
	}
	// 调用闭包函数
	r := sumFunc()
	fmt.Println(r)

	// 有入参的闭包函数
	sumFunc2 := func(n int) int {
		cur := 0
		for i := 1; i <= n; i++ {
			cur += i
		}
		return cur
	}
	r = sumFunc2(200)
	fmt.Println(r)

	// 使用闭包定义 lambda 回调函数入参
	strA := "I have had my invitation to this world's festival"
	strB := " i have had my invitation to this world's festival"
	isEqual := myCompare(strA, strB, func(a, b string) bool {
		a = strings.ToLower(strings.TrimSpace(a))
		b = strings.ToLower(strings.TrimSpace(b))
		return a == b
	})
	fmt.Println(isEqual)

}

func myCompare(a, b string, equal func(a, b string) bool) bool {
	return equal(a, b)
}
