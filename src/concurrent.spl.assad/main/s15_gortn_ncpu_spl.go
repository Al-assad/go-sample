package main

import (
	"fmt"
	"runtime"
)

// @desc GOMAXPROCS 设置多核并行计算

// 最高可用核心数
const NCPU = 6

func main() {
	// 设置最高可用核心数
	runtime.GOMAXPROCS(NCPU) // 会返回系统可用全部核心数
	doAll()
}

// 多核心并行运行
func doAll() {
	sem := make(chan int, NCPU)
	// 设置合理 sem 缓冲区，一般和 NCPU 一致，完全使用过全部 NCPU 资源函
	for i := 0; i < NCPU; i++ {
		go doPart(sem)
	}
	// 等待 NCPU 任务完成，释放通道 sem 的缓冲区
	for i := 0; i < NCPU; i++ {
		<-sem // 等待一个任务完成
	}

}

func doPart(sem chan int) {
	// do something
	fmt.Println(fibonacci2(40))
	sem <- 1
}

func fibonacci2(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	return res
}
