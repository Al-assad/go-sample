package main

import (
	"fmt"
	"time"
)

// 递归函数示例， fibonacci 计算

func main() {
	start := time.Now()

	val := 0
	// 获取前 40 个 fibonacci
	for i := 0; i <= 40; i++ {
		val = fibonacci(i)
		fmt.Printf("fibonacci(%d)=%d \n", i, val)
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}
