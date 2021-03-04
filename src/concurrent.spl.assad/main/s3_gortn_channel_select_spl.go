package main

import (
	"fmt"
	"time"
)

// @desc channel 通道选择器示例
// go select 用于监听 io 操作，当 io 操作发生时，触发对应动作
/*
select 做的就是：选择处理列出的多个通信情况中的一个。
 	- 如果都阻塞了，会等待直到其中一个可以处理
	- 如果多个可以处理，随机选择一个
	- 如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）
*/

func main() {
	//selectTest1()
	selectTest2()
}

// select 测试1
func selectTest1() {
	ch1, ch2 := make(chan string), make(chan string)
	go func() {
		ch1 <- "hello"
		close(ch1)
	}()

	go func() {
		ch2 <- "hi"
		close(ch2)
	}()

	// 可能输出 ch1、也可能输出 ch2，取决于哪一个协程先向 channel 发送数据
	select {
	case v1 := <-ch1:
		fmt.Println("receive ch1:", v1)
	case v2 := <-ch2:
		fmt.Println("receive ch2:", v2)
	}
}

// select 示例2：模拟消费者消费多个生产者
func selectTest2() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go produceA(ch1)
	go produceB(ch2)
	consume(ch1, ch2)
}

// 生产者 A
func produceA(ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("PA-%d", i)
		time.Sleep(1 * time.Second)
	}
}

// 生产者 B
func produceB(ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("PB-%d", i)
		time.Sleep(1 * time.Second)
	}
}

func consume(ch1, ch2 chan string) {
	for {
		select {
		case pa := <-ch1:
			fmt.Println("receive:", pa)
		case pb := <-ch2:
			fmt.Println("receive", pb)
		}
	}
}
