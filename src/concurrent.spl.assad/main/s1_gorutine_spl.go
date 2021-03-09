package main

import (
	"fmt"
	"time"
)

// @desc go 协程基本使用

func main() {
	goroutineSpl1()
	/**
	main start
	task2 start
	task1 start
	task2 completed
	task1 completed
	main completed
	*/
	goroutineSpl2()
}

// 协程使用示例
func goroutineSpl1() {
	fmt.Println("main start")
	go task1() // 运行协程 task1
	go task2() // 运行协程 task2
	time.Sleep(10 * time.Second)
	fmt.Println("main completed")
}

func task1() {
	fmt.Println("task1 start")
	time.Sleep(5 * time.Second)
	fmt.Println("task1 completed")
}

func task2() {
	fmt.Println("task2 start")
	time.Sleep(2 * time.Second)
	//runtime.Goexit()  终止当前协程
	fmt.Println("task2 completed")
}

// 协程使用示例2：主线程等待协程执行结束
func goroutineSpl2() {
	ch := make(chan bool)
	fmt.Println("main start")
	go func() {
		fmt.Println("do something")
		ch <- true
	}()
	<-ch
	fmt.Println("main completed")
}
