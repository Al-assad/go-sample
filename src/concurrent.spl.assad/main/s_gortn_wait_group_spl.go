package main

import (
	"fmt"
	"sync"
)

// @desc 使用 waitGroup 优雅等待所有协程执行结束

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("main start")
	go doSomething(wg.Done)
	go doSomething(wg.Done)

	wg.Wait()
	fmt.Println("main finish")

	/*
		main start
		do something in goroutine
		do something in goroutine
		main finish
	*/
}

func doSomething(done func()) {
	defer done()
	fmt.Println("do something in goroutine")
}

// 也可以使用空置的 select{} 实现永久等待
//func runForever()  {
//	go doSomething()
//	select{}
//}
