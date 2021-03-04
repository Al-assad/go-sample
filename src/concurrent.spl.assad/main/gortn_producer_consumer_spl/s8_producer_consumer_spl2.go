package main

import (
	"fmt"
	"math/rand"
	"time"
)

// @desc go 协程实现生产者-消费者示例：多生产者、单消费者
/*
produceA： 0
produceB： 0
consume from produceA 0
consume from produceB 0
produceB： 1
consume from produceB 1
produceB： 2
produceA： 1
consume from produceB 2
consume from produceA 1
produceA： 2
produceA： 3
consume from produceA 2
consume from produceA 3
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go produceA(ch1)
	go produceB(ch2)
	go consumerS(ch1, ch2)
	select {}
}

// 生产者 A
func produceA(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println("produceA：", i)
		ch <- i
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func produceB(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println("produceB：", i)
		ch <- i
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

// 消费者，轮流消费生产者
func consumerS(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("consume from produceA", v)
		case v := <-ch2:
			fmt.Println("consume from produceB", v)
		}
	}
}
