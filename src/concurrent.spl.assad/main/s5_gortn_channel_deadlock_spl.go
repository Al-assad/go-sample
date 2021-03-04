package main

import "fmt"

// @desc go 协程 channel 死锁示例

func main() {
	//deadLockSpl()
	//deLockSpl()x
}

// 死锁演示
func deadLockSpl() {
	ch := make(chan int)
	ch <- 1
	a := <-ch
	fmt.Println(a)
	/**
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send]:
	main.deadLockSpl()
	        /Users/yulin/Workplace/go-sample/src/concurrent.spl.assad/main/s2_gortncomm_channel_detail_spl.go:100 +0x59
	main.main()
	        /Users/yulin/Workplace/go-sample/src/concurrent.spl.assad/main/s2_gortncomm_channel_detail_spl.go:24 +0x25
	*/
}

// deadLockSpl 避免死锁
func deLockSpl() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	a := <-ch
	fmt.Println(a)
	/**
	1
	*/
}
