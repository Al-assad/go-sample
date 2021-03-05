package main

import (
	"fmt"
	"time"
)

// @desc 定时器 Ticker 和 Channel 的结合使用（延迟、超时定时器）

func main() {
	delayTickBoom()
}

// 使用 Channel 实现延迟终止的 Ticker
// boom 定时器作为 tick 的延迟终止指令发送定时器
func delayTickBoom() {
	tick := time.Tick(1e9)
	boom := time.After(5e9)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			time.Sleep(5e8)
		}
	}
	/*
		tick.
		tick.
		tick.
		tick.
		boom!
	*/
}

func simpleTimeout() {

}
