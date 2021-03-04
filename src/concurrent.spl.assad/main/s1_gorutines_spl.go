package main

import (
	"fmt"
	"time"
)

// @desc go 协程基本使用

func main() {
	fmt.Println("main start")
	go task1()
	go task2()
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
	fmt.Println("task2 completed")
}
