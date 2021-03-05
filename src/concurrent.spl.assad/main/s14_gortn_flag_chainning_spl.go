package main

import (
	"flag"
	"fmt"
)

// @desc 通过 flag 快速生成大量协程

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost

	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0

	// start the chaining 开始链接
	x := <-leftmost // wait for completion 等待完成

	fmt.Println(x)

	// 结果： 100000 ， 不到 200ms
}
