package main

import (
	"fmt"
	"time"
)

// @desc go 协程实现生产者-消费者示例：单生产者、单消费者
/*
	produce： 0
	consume: 0
	produce： 1
	consume: 1
	produce： 2
	consume: 2
	produce： 3
	consume: 3
	produce： 4
	consume: 4
	produce： 5
	consume: 5
	produce： 6
	consume: 6
*/
func main() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	select {}
}

// 生产者
func produce(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println("produce：", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}
}

// 消费者
func consumer(ch chan int) {
	for {
		data := <-ch
		fmt.Println("consume:", data)
	}
}
