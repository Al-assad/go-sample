package main

import (
	"fmt"
	"time"
)

// @desc go 协程基本使用

func main() {
	fmt.Println("main start")
	go task1() // 运行协程 task1
	go task2() // 运行协程 task2
	time.Sleep(10 * time.Second)
	fmt.Println("main completed")
	/**
	main start
	task2 start
	task1 start
	task2 completed
	task1 completed
	main completed
	*/
}

func task1() {
	fmt.Println("task1 start")
	time.Sleep(5 * time.Second)
	fmt.Println("task1 completed")
}

func task2() {
	fmt.Println("task2 start")
	time.Sleep(2 * time.Second)
	fmt.Println("task2 completed")
}
