package main

import (
	"fmt"
	"sync"
)

// @desc go 协程共享数据竞争问题：使用 channel、lock 方式

func main() {
	dataCompetitionSpl()
}

var num = 0

// 数据竞争的例子
func dataCompetitionSpl() {
	var wg sync.WaitGroup
	wg.Add(2)

	go updateNum(10)
	go updateNum(10)

	wg.Wait()
	fmt.Println(num)
}

func updateNum(update int) {
	num += update
}

//// 使用协程的方式
//func withChannelSpl() {
//
//}
//
//func updateNum2(num int, ch chan int) {
//
//}
