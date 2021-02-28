package main

import (
	"fmt"
	"time"
)

// 递归函数示例， fibonacci 计算，使用内存缓存

var fibs [50]uint64

func main() {
	start := time.Now()

	val := uint64(0)
	// 获取前 40 个 fibonacci
	for i := 0; i <= 40; i++ {
		val = fibonacci2(i)
		fmt.Printf("fibonacci(%d)=%d \n", i, val)
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func fibonacci2(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return res
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return res
}
