package main

import (
	"fmt"
	"math/rand"
	"time"
)

// @desc 定时器 Ticker 和 Channel 的结合使用（延迟、超时定时器）

func main() {
	//tickBoom()
	//timeoutSpl()
	timeoutSpl2()
}

// 使用 Channel 实现延迟终止的 Ticker
// boom 定时器作为 tick 的延迟终止指令发送定时器
func tickBoom() {
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

// 实现执行超时限制，简单实现
func timeoutSpl() {
	// 超时消息发送协程
	timeout := make(chan bool, 1)
	go func(timeoutSec int) {
		time.Sleep(time.Duration(timeoutSec) * time.Second)
		timeout <- true
	}(3)
	// 需要被限制执行超时的协程
	// 如果没有在 3s 内执行结束，会被强制终止
	ch := make(chan string)
	go func() {
		runSec := time.Duration(rand.Intn(10)) * time.Second
		time.Sleep(runSec)
		ch <- "runSec: " + runSec.String()
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-timeout:
		fmt.Println("timeout")
		break
	}
}

// 实现超时限制，使用 time.After
func timeoutSpl2() {
	timeoutSec := time.Duration(3) * time.Second

	// 需要被限制执行超时的协程
	ch := make(chan string)
	go func() {
		runSec := time.Duration(rand.Intn(10)) * time.Second
		time.Sleep(runSec)
		ch <- "runSec: " + runSec.String()
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(timeoutSec):
		fmt.Println("timeout")
		break
	}

}
