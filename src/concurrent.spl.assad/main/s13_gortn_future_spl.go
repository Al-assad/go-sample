package main

import "fmt"

// @desc 实现 feature 模式

func main() {
	fmt.Println("start inverse fibonacci")
	future := inverseFibo(20)

	fmt.Println("do other thing.")
	fmt.Println("do other thing..")

	fmt.Println("get fibonacci val", <-future)
	fmt.Println("after inverse fibonacci")

	/*
		start inverse fibonacci
		do other thing.
		do other thing..
		get fibonacci val 10946
		after inverse fibonacci
	*/
}

// 计算 fibonacci，返回 Future
func inverseFibo(n int) chan int {
	future := make(chan int)
	go func() {
		future <- myFibonacci(n)
	}()
	return future
}

func myFibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = myFibonacci(n-1) + myFibonacci(n-2)
	}
	return res
}
