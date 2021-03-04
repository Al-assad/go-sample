package main

import (
	"fmt"
	"math/rand"
	"time"
)

// @desc go 协程实现生产者-消费者示例：单生产者、多消费者
/*
produce 0
consumeB 0
produce 1
consumeA 1
produce 2
consumeB 2
produce 3
consumeA 3
produce 4
consumeB 4
produce 5
produce 6
consumeB 6
*/

func main() {
	ch := make(chan int)
	go produceS(ch)
	go consumeA(ch)
	go consumeB(ch)
	select {}
}

func produceS(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println("produce", i)
		ch <- i
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func consumeA(ch chan int) {
	for {
		data := <-ch
		fmt.Println("consumeA", data)
	}
}
func consumeB(ch chan int) {
	for {
		data := <-ch
		fmt.Println("consumeB", data)
	}
}
