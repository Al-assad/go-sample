package main

import (
	"fmt"
	"time"
)

// @desc go ticker 定时器使用示例

func main() {
	//simpleTicker()
	//simpleTimer()
	myTimer(2e9, 1e9, 10e9, func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})
}

// Ticker 定时器使用
// 延迟定时 1s 执行
func simpleTicker() {
	ticker := time.NewTicker(1 * time.Second) // 创建定时器，也可以通过 time.Ticker 创建
	defer ticker.Stop()                       // 停止定时器
	// 定时器动作
	for range ticker.C {
		fmt.Println("tick.")
	}
	/*
		tick.
		tick.
		tick.
		tick.
		tick.
		tick.
		tick.
	*/
}

// Timer 延迟器使用
// 演示延迟 5s 后执行
func simpleTimer() {
	timer := time.NewTimer(5 * time.Second) // 和 time.After 类似
	<-timer.C
	fmt.Println("Boom!")
}

// 实现一个简单的支持延迟启动的定时函数（仅仅延迟 time. 定时函数的相关使用）
// delay 延迟启动时间
// interval 运行间隔时间
// timeout 总运行超时时间
// action 运行动作
func myTimer(delay time.Duration, interval time.Duration, timeout time.Duration, action func()) {
	// 延迟启动
	delayT := time.NewTimer(delay)
	<-delayT.C
	delayT.Stop()
	// 超时终止
	timeoutT := time.NewTimer(timeout)
	// 定时运行
	tickT := time.NewTicker(interval)
	for {
		select {
		case <-tickT.C:
			action()
		case <-timeoutT.C:
			return
		}
	}

}
