package main

import (
	"fmt"
)

// @desc panic 运行时异常的抛出、捕获处理
/*
panic 的使用规范：
	1. 在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic ()
	2. 向包的调用者返回 error（而不是 panic）
*/

func main() {
	//badCall()
	catchPanicFromBadCall()
}

// 抛出 panic 异常
// panic 异常强制终止当前程序，并打印出异常信息
func badCall() {
	fmt.Println("Before throw panic")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("After throw panic")

	/*
		Before throw panic
		panic: A severe error occurred: stopping the program!

		goroutine 1 [running]:
		main.badCall()
		        /Users/yulin/Workplace/go-sample/src/spl.assad/main/h1_error_panic.go:17 +0x95
		main.main()
		        /Users/yulin/Workplace/go-sample/src/spl.assad/main/h1_error_panic.go:10 +0x25
	*/

}

// 捕获处理 panic 异常，类似 Java Catch
// recover() 只能在 defer 中使用，
//panic 抛出后会被后面堆栈的 defer 中的 recover() 中捕获，如果没有 recover, 则会传递到 os.Exit(1)
func catchPanicFromBadCall() {
	// 捕获处理 panic 异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Receive runtime panic: %v \n", err)
		}
	}()

	fmt.Println("Before Call badCall()")
	badCall()
	fmt.Println("After Call badCall()")

	/*
		Before Call badCall()
		Before throw panic
		Receive runtime panic: A severe error occurred: stopping the program!
	*/
}
