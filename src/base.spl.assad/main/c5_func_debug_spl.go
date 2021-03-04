package main

import (
	"base.spl.assad/func_foo"
	"fmt"
	"log"
	"runtime"
	"time"
)

// @desc 函数调试, 打印函数执行位置，统计函数执行时间

func main() {

	// 打印函数具体位置
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	func_foo.Minus(1, 2)
	where()
	// 2021/02/27 12:05:38 /Users/yulin/Workplace/test-go/src/spl.assad/main/c5_func_debug_spl.g2o:17
	// 2021/02/27 12:05:38 /Users/yulin/Workplace/test-go/src/spl.assad/main/c5_func_debug_spl.go:19

	// 打印函数具体位置，使用 log 函数
	log.SetFlags(log.Ldate + log.Ltime + log.Llongfile)
	where2 := log.Print
	where2()
	func_foo.Minus(1, 2)
	where2()
	// 2021/02/27 12:05:38 /Users/yulin/Workplace/test-go/src/spl.assad/main/c5_func_debug_spl.go:24:
	// 2021/02/27 12:05:38 /Users/yulin/Workplace/test-go/src/spl.assad/main/c5_func_debug_spl.go:26:

	// 统计函数执行时间
	start := time.Now()
	longCompute()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)

}

func longCompute() {
	fmt.Println("start longCompute()")
	time.Sleep(time.Second * 2)
	fmt.Println("finish longCompute()")
}
