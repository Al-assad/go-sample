package main

import (
	"fmt"
	"time"
)

// @desc 协程 recover 示例
// 一个用到 recover 的程序停掉服务器内部一个失败的协程，并不会影响其他协程的工作

func main() {
	go safelyDo(1)
	go safelyDo(2)
	go safelyDo(4)
	go safelyDo(6)

	time.Sleep(10e9)
	/*
		do successful 2
		do fail 1
		do successful 6
		do successful 4
	*/
}

func safelyDo(val int) {
	// panic 温和处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("do fail", val)
		}
	}()
	do(val)
}

func do(val int) {
	if val%2 == 1 {
		panic("throw panic！")
	} else {
		fmt.Println("do successful", val)
	}
}
